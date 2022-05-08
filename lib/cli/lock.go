package cli

import (
	"github.com/spf13/cobra"
	"github.com/tlopo-go/secrets/lib/app"
)

var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Locks the secrets database",
	Long:  "",
	Run:   lockRun,
}

func lockRun(cmd *cobra.Command, args []string) {
	app.LockDB()
}
