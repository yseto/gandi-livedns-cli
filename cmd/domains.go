package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var GetDomains = cli.Command{
	Before: Before,
	Name:   "domains",
	Usage:  "list domains",
	Action: func(cCtx *cli.Context) error {
		d, err := client.GetDomains()
		if err != nil {
			return err
		}
		for i := range d {
			fmt.Println(d[i].FQDN)
		}
		return nil
	},
}
