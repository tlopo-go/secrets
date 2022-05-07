package cli

import (
	"github.com/spf13/cobra"
	"github.com/tlopo-go/secrets/lib/app"
	"log"
)

var unlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Unlocks the secrets database",
	Long:  "",
	Run:   unlockRun,
}

func unlockRun(cmd *cobra.Command, args []string) {
	if !app.IsDBLocked() {
		log.Println("Database is already unlocked")
	}
	app.UnlockDB()
}
