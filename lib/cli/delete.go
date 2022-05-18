package cli

import (
	"github.com/spf13/cobra"
	"github.com/tlopo-go/secrets/lib/app"
	k "github.com/tlopo-go/secrets/lib/keepass"
	"log"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a secret",
	Long:  "",
	Run:   deleteRun,
}

type deleteCmdArgs struct {
	service string
}

var deleteArgs deleteCmdArgs

func init() {
	deleteCmd.Flags().StringVarP(&deleteArgs.service, "service", "s", "", "service name")
	deleteCmd.MarkFlagRequired("service")
}

func deleteRun(cmd *cobra.Command, args []string) {
	app.ValidateUnlocked()
	kp := k.KeePass{app.GetDatabasePath(), app.GetMasterPassword()}
	err := kp.Delete(deleteArgs.service)
	if err != nil {
		log.Fatal(err)
	}
}
