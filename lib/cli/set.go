package cli

import (
	"github.com/spf13/cobra"
	"github.com/tlopo-go/secrets/lib/app"
	k "github.com/tlopo-go/secrets/lib/keepass"
	"log"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret",
	Long:  "",
	Run:   setRun,
}

type setCmdArgs struct {
	service  string
	account  string
	password string
}

var setArgs setCmdArgs

func init() {
	setCmd.Flags().StringVarP(&setArgs.service, "service", "s", "", "service name")
	setCmd.Flags().StringVarP(&setArgs.account, "account", "a", "", "account")
	setCmd.Flags().StringVarP(&setArgs.password, "password", "p", "", "password")

	for _, str := range []string{"service", "account", "password"} {
		setCmd.MarkFlagRequired(str)
	}
}

func setRun(cmd *cobra.Command, args []string) {
	app.ValidateUnlocked()
	kp := k.KeePass{app.GetDatabasePath(), app.GetMasterPassword()}
	err := kp.Write(k.Secret{setArgs.service, setArgs.account, setArgs.password})
	if err != nil {
		log.Fatal(err)
	}
}
