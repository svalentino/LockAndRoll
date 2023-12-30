package backend

type Backend interface {
	// Operations on secrets
	ListSecrets() ([]string, error)
	CreateSecret(secret Secret) error
	GetSecret(name string) (*Secret, error)
	UpdateSecret(secret Secret) error
	DeleteSecret(name string) (*Secret, error)
}

type backendCatalogue map[string]Backend

var BackendDispatcher backendCatalogue

func init() {
	// Register all implemented backends here
	BackendDispatcher = backendCatalogue{}
	BackendDispatcher["macos-keychain"] = NewBackendMacosKeychain()
}
