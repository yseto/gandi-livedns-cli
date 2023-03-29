package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Snapshot = cli.Command{
	Before: Before,
	Name:   "snapshot",
	Usage:  "operate of snapshot",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "domain",
			Usage:    "domain name",
			Required: true,
		},
	},
	Subcommands: []*cli.Command{
		{
			Name:  "list",
			Usage: "list of snapshots",
			Action: func(cCtx *cli.Context) error {
				domainName := cCtx.String("domain")
				d, err := client.GetSnapshots(domainName)
				if err != nil {
					return err
				}
				fmt.Println("ID\tName\tCreatedAt\tAutomatic")
				for _, v := range d {
					fmt.Println(v)
				}
				return nil
			},
		},
		{
			Name:  "show",
			Usage: "detail of snapshot",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "id",
					Usage:    "snapshot id",
					Required: true,
				},
			},
			Action: func(cCtx *cli.Context) error {
				domainName := cCtx.String("domain")
				id := cCtx.String("id")
				d, err := client.GetSnapshot(domainName, id)
				if err != nil {
					return err
				}
				fmt.Printf("$ORIGIN %s.\n", domainName)
				for _, v := range d {
					fmt.Println(v)
				}
				return nil
			},
		},
		{
			Name:  "delete",
			Usage: "delete a snapshot",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "id",
					Usage:    "snapshot id",
					Required: true,
				},
			},
			Action: func(cCtx *cli.Context) error {
				domainName := cCtx.String("domain")
				id := cCtx.String("id")
				message, err := client.DeleteSnapshot(domainName, id)
				if err != nil {
					return err
				}
				fmt.Println(*message)
				return nil
			},
		},
		{
			Name:  "create",
			Usage: "create a snapshot",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "name",
					Usage:    "snapshot name",
					Required: true,
				},
			},
			Action: func(cCtx *cli.Context) error {
				domainName := cCtx.String("domain")
				name := cCtx.String("name")
				message, err := client.CreateSnapshot(domainName, name)
				if err != nil {
					return err
				}
				fmt.Println(*message)
				return nil
			},
		},
		{
			Name:  "update",
			Usage: "update a snapshot",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "id",
					Usage:    "snapshot id",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "name",
					Usage:    "snapshot name",
					Required: true,
				},
			},
			Action: func(cCtx *cli.Context) error {
				domainName := cCtx.String("domain")
				name := cCtx.String("name")
				id := cCtx.String("id")
				message, err := client.UpdateSnapshot(domainName, id, name)
				if err != nil {
					return err
				}
				fmt.Println(*message)
				return nil
			},
		},
	},
}
