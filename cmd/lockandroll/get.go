package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var export bool
var secretName string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <secret name>",
	Short: "Retrieve the value of a secret",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		secret, err := b.GetSecret(args[0])
		if err != nil {
			panic(err)
		}
		for k, v := range secret.Variables {
			if export {
				fmt.Printf("export ")
			}
			fmt.Printf("%s=%q\n", k, v)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().BoolVarP(&export, "export", "e", false, "Easily export the secret as an environment variable(s)")
}
