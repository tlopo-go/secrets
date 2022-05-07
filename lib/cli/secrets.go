package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	SecretsCmd.AddCommand(getCmd)
	SecretsCmd.AddCommand(setCmd)
	SecretsCmd.AddCommand(unlockCmd)
}

var SecretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Command line secret manager",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
