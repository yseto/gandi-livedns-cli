package main

import (
	"fmt"

	"github.com/yseto/gandi-livedns-cli/config"

	"github.com/urfave/cli/v2"
)

var ApiKey = cli.Command{
	Name:  "apikey",
	Usage: "api key utility",
	Subcommands: []*cli.Command{
		{
			Name:  "save",
			Usage: "save apikey",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "apikey",
					Usage:    "apikey",
					Required: true,
				},
			},
			Action: func(cCtx *cli.Context) error {
				apiKey := cCtx.String("apikey")

				path, err := config.SaveApiKey(apiKey)
				if err != nil {
					return err
				}
				fmt.Printf("Saved: %s\n", path)

				return nil
			},
		},
		{
			Name:  "show",
			Usage: "show apikey",
			Action: func(cCtx *cli.Context) error {
				key, err := config.LoadApiKey()
				if err != nil {
					fmt.Println("missing API KEY")
					return nil
				}
				fmt.Printf("API KEY : %s\n", key)
				return nil
			},
		},
		{
			Name:  "delete",
			Usage: "delete config file",
			Action: func(cCtx *cli.Context) error {
				err := config.DeleteApiKey()
				if err != nil {
					return err
				}
				fmt.Println("deleted config file")
				return nil
			},
		},
	},
}
