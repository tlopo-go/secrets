package cli

import (
	"github.com/spf13/cobra"
	"github.com/tlopo-go/secrets/lib/app"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the secrets database",
	Long:  "",
	Run:   initRun,
}

func initRun(cmd *cobra.Command, args []string) {
	app.InitDB()
}
