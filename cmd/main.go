package main

import (
	"fmt"
	"os"

	"github.com/yseto/gandi-livedns-cli/config"
	"github.com/yseto/gandi-livedns-cli/gandi"

	"github.com/urfave/cli/v2"
)

var client *gandi.Client

var Before = func(cCtx *cli.Context) error {
	apikey, err := config.LoadApiKey()
	if err != nil {
		return err
	}
	client = gandi.NewClient(apikey)
	return nil
}

func main() {
	app := &cli.App{
		Usage: "Gandi LiveDNS cli",
		Commands: []*cli.Command{
			&GetDomains,
			&Export,
			&CreateRecord,
			&DeleteRecord,
			&Snapshot,
			&ApiKey,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
