/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/svalentino/lockandroll/pkg/backend"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update <secret name>",
	Short: "Update the value of a secret",
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

		err = b.UpdateSecret(backend.Secret{
			Name:      name,
			Variables: vs,
		})
		if err != nil {
			fmt.Printf("Error updating secret '%s': %v\n", name, err)
		}
		fmt.Printf("Secret '%s' updated successfully.\n", name)

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
