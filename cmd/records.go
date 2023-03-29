package main

import (
	"errors"
	"fmt"

	"github.com/yseto/gandi-livedns-cli/gandi"

	"github.com/urfave/cli/v2"
)

var Export = cli.Command{
	Before: Before,
	Name:   "export",
	Usage:  "export zone",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "domain",
			Usage:    "domain name",
			Required: true,
		},
	},
	Action: func(cCtx *cli.Context) error {
		domainName := cCtx.String("domain")
		d, err := client.GetRecords(domainName)
		if err != nil {
			return err
		}
		fmt.Printf("$ORIGIN %s.\n", domainName)
		for _, v := range d {
			fmt.Println(v)
		}
		return nil
	},
}

var ErrMissingRecord = errors.New("need Record")

var CreateRecord = cli.Command{
	Before: Before,
	Name:   "rrcreate",
	Usage:  "create record",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "domain",
			Usage:    "domain name",
			Required: true,
		},
		&cli.BoolFlag{
			Name:  "replace",
			Usage: "replace a record",
		},
	},
	Action: func(cCtx *cli.Context) error {
		domainName := cCtx.String("domain")
		replace := cCtx.Bool("replace")
		argRecords := cCtx.Args().Slice()

		var records []gandi.Record
		for i := range argRecords {
			v, err := gandi.ParseRecord(argRecords[i])
			if err != nil {
				return err
			}
			records = append(records, *v)
		}

		var m gandi.RecordMerger
		for i := range records {
			if err := m.Merge(records[i]); err != nil {
				return err
			}
		}
		records = m.Output()
		if len(records) == 0 {
			return ErrMissingRecord
		}

		if replace {
			rrRecords := m.ReplaceOutput()
			for i := range rrRecords {
				message, err := client.ReplaceRecord(domainName, rrRecords[i])
				if err != nil {
					return err
				}
				fmt.Println(*message)
			}
			return nil
		}

		for i := range records {
			message, err := client.CreateRecord(domainName, records[i])
			if err != nil {
				return err
			}
			fmt.Println(*message)
		}
		return nil
	},
}

var DeleteRecord = cli.Command{
	Before: Before,
	Name:   "rrdelete",
	Usage:  "delete record",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "domain",
			Usage:    "domain name",
			Required: true,
		},
	},
	Action: func(cCtx *cli.Context) error {
		domainName := cCtx.String("domain")
		argRecords := cCtx.Args().Slice()

		v, err := gandi.ParseRecordWithoutValue(argRecords)
		if err != nil {
			return err
		}

		message, err := client.DeleteRecord(domainName, *v)
		if err != nil {
			return err
		}
		fmt.Println(*message)

		return nil
	},
}
