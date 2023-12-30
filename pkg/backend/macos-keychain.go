package backend

import (
	"fmt"
	"strings"

	"github.com/keybase/go-keychain"
)

type BackendMacosKeychain struct {
	accountName string
	namePrefix  string
	class       keychain.SecClass
}

func NewBackendMacosKeychain() Backend {
	return &BackendMacosKeychain{
		accountName: "LockAndRoll",
		namePrefix:  "lockandroll",
		class:       keychain.SecClassGenericPassword,
	}
}

func (b *BackendMacosKeychain) newKeychainItem() keychain.Item {
	item := keychain.NewItem()
	item.SetSecClass(b.class)
	item.SetAccount(b.accountName)
	return item
}

func (b *BackendMacosKeychain) stripPrefix(name string) string {
	if strings.HasPrefix(name, b.namePrefix+"-") {
		return name[len(b.namePrefix)+1:]
	}
	return name
}

func (b *BackendMacosKeychain) addPrefix(name string) string {
	if !strings.HasPrefix(name, b.namePrefix+"-") {
		return fmt.Sprintf("%s-%s", b.namePrefix, name)
	}
	return name
}

func (b *BackendMacosKeychain) secretExists(name string) (bool, error) {
	secrets, err := b.ListSecrets()
	if err != nil {
		return false, err
	}

	name = b.stripPrefix(name)
	for _, secret := range secrets {
		if secret == name {
			return true, nil
		}
	}

	return false, nil
}

func (b *BackendMacosKeychain) ListSecrets() ([]string, error) {
	query := b.newKeychainItem()
	query.SetMatchLimit(keychain.MatchLimitAll)
	query.SetReturnAttributes(true)

	results, err := keychain.QueryItem(query)
	if err != nil {
		return nil, fmt.Errorf("error in querying the backend: %w", err)
	}

	output := []string{}
	for _, r := range results {
		output = append(output, b.stripPrefix(r.Service))
	}

	return output, nil
}

func (b *BackendMacosKeychain) CreateSecret(secret Secret) error {
	exists, err := b.secretExists(secret.Name)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("secret already exists: %q", secret.Name)
	}

	// Transform secret.Variables into a JSON string
	js, err := secret.Variables.ToJSON()
	if err != nil {
		return err
	}

	// Prepare item to be saved into Keychain
	item := b.newKeychainItem()
	item.SetService(b.addPrefix(secret.Name))
	item.SetData([]byte(js))
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleDefault)

	// Save
	if err = keychain.AddItem(item); err != nil {
		return fmt.Errorf("error in saving the secret: q: %w", secret.Name, err)
	}

	return nil
}

func (b *BackendMacosKeychain) GetSecret(name string) (*Secret, error) {
	// Prepare query
	query := b.newKeychainItem()
	query.SetService(b.addPrefix(name))
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnData(true)

	// Query
	results, err := keychain.QueryItem(query)
	if err != nil {
		return nil, fmt.Errorf("error in querying the backend: %w", err)
	}

	// Check if the secret exists
	if len(results) == 0 {
		return nil, fmt.Errorf("secret doesn't exists: %q", name)
	}

	// Decode the secret
	secret := Secret{
		Name: name,
	}
	if err := secret.Variables.FromJSON(string(results[0].Data)); err != nil {
		return nil, fmt.Errorf("error in decoding the secret: %w", err)
	}

	return &secret, nil
}

func (b *BackendMacosKeychain) UpdateSecret(secret Secret) error {
	exists, err := b.secretExists(secret.Name)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("secret doesn't exists: %q", secret.Name)
	}

	// Transform secret.Variables into a JSON string
	js, err := secret.Variables.ToJSON()
	if err != nil {
		return err
	}

	// Prepare item to be saved into Keychain
	item := b.newKeychainItem()
	item.SetService(b.addPrefix(secret.Name))
	item.SetData([]byte(js))
	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleAlways)

	// Save
	if err = keychain.UpdateItem(item, item); err != nil {
		return fmt.Errorf("error in saving the secret: %q: %w", secret.Name, err)
	}

	return nil
}

func (b *BackendMacosKeychain) DeleteSecret(name string) (*Secret, error) {
	// Get backup copy of the secret
	secret, err := b.GetSecret(name)
	if err != nil {
		return nil, err
	}

	// Prepare query
	query := b.newKeychainItem()
	query.SetService(b.addPrefix(name))
	query.SetMatchLimit(keychain.MatchLimitOne)

	// Delete
	if err := keychain.DeleteItem(query); err != nil {
		return nil, fmt.Errorf("error in deleting the secret: %q: %w", name, err)
	}

	return secret, nil
}
