package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all stored secrets",
	Run: func(cmd *cobra.Command, args []string) {
		secrets, err := b.ListSecrets()
		if err != nil {
			panic(err)
		}
		for _, secret := range secrets {
			fmt.Println(secret)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
