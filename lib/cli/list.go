package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tlopo-go/secrets/lib/app"
	k "github.com/tlopo-go/secrets/lib/keepass"
	"log"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List secrets",
	Long:  "",
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	app.ValidateUnlocked()
	kp := k.KeePass{app.GetDatabasePath(), app.GetMasterPassword()}
	names, err := kp.List()
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range names {
		fmt.Println(name)
	}
}
