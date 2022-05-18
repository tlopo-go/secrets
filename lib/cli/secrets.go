package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	SecretsCmd.AddCommand(getCmd)
	SecretsCmd.AddCommand(setCmd)
	SecretsCmd.AddCommand(deleteCmd)
	SecretsCmd.AddCommand(listCmd)
	SecretsCmd.AddCommand(unlockCmd)
	SecretsCmd.AddCommand(lockCmd)
	SecretsCmd.AddCommand(initCmd)
}

var SecretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Command line secret manager",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
