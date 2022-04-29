package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	k "github.com/tlopo-go/secrets/lib/keepass"
	"os"
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
	kp := k.KeePass{"/tmp/db.kdbx", "1234"}
	s, err := kp.Read(service)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%#v\n", s)
}
