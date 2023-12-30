package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/svalentino/lockandroll/pkg/backend"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <secret name> [secret value (must be JSON)]",
	Short: "Create and save new secret",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		value := ""

		// If 2nd positional argument is provided, take it as value of the secret
		// Otherwise, prompt the user for the value
		if len(args) > 1 {
			value = args[1]
		} else {
			value = PromptSecretValue()
		}

		// Parse value of the secret as JSON
		vs := backend.Variables{}
		err := vs.FromJSON(value)
		if err != nil {
			fmt.Printf("The value provided for the secret is not valid JSON: %v\n", err)
			os.Exit(1)
		}

		// Store secret into the backend
		err = b.CreateSecret(backend.Secret{
			Name:      name,
			Variables: vs,
		})
		if err != nil {
			fmt.Printf("Error creating secret '%s': %v\n", name, err)
			os.Exit(1)
		}
		fmt.Printf("Secret '%s' created successfully.\n", name)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

}
