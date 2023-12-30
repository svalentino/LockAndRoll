package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var showDeleted bool
var deleteAll bool

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [secret name]",
	Short: "Delete a secret",
	Args: func(cmd *cobra.Command, args []string) error {
		// Fail if both --all and a secret name are specified
		if err := cobra.NoArgs(cmd, args); err != nil && deleteAll {
			return fmt.Errorf("cannot specify both '--all' and a 'secret name'")
		}
		// Fail if --all is not specified but no secret name was provided
		if err := cobra.ExactArgs(1)(cmd, args); err != nil && !deleteAll {
			return fmt.Errorf("requires a secret name argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		toDelete := []string{}
		if deleteAll {
			toDelete, _ = b.ListSecrets()
		} else {
			toDelete = append(toDelete, args[0])
		}

		for _, name := range toDelete {
			backupSecret, err := b.DeleteSecret(name)
			if err != nil {
				fmt.Printf("Error deleting secret '%s': %v\n", name, err)
				os.Exit(1)
			}
			fmt.Printf("Secret '%s' deleted successfully.\n", name)

			if showDeleted {
				fmt.Printf("Secret '%s' content:\n", name)
				for k, v := range backupSecret.Variables {
					fmt.Printf("%s=%q\n", k, v)
				}
				fmt.Println()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().BoolVarP(&showDeleted, "show", "s", false, "Show the full content of the deleted secret")
	deleteCmd.Flags().BoolVar(&deleteAll, "all", false, "Delete all secrets")
}
