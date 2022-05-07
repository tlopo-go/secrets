package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tlopo-go/secrets/lib/app"
	k "github.com/tlopo-go/secrets/lib/keepass"
	"log"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves a secret",
	Long:  "",
	Run:   run,
}

var service string
var user string
var pass string

func init() {
	getCmd.Flags().StringVarP(&service, "service", "s", "", "service name")
	getCmd.MarkFlagRequired("service")
}

func run(cmd *cobra.Command, args []string) {
	if app.IsDBLocked() {
		log.Fatal("Database is locked")
	}
	kp := k.KeePass{app.GetDatabasePath(), app.GetMasterPassword()}
	s, err := kp.Read(service)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s.ToJson())
}
