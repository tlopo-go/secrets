package main

import (
	"github.com/tlopo-go/secrets/lib/cli"
	"os"
)

func main() {
	if err := cli.SecretsCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
