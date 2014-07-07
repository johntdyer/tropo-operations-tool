package main

import (
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	guid,
	list,
	lookup,
}

func assert(err error) {
	if err != nil {
		logger.Fatal(err)
	}
}

// tropo lookup address 4075551212 --user --application
// tropo lookup address 99999999   --user --application
// tropo lookup address abc123     --user --application
// tropo lookup appplication 1234  --addresses

// tropo guid abc123
// tropo check dns 4075551212

// tropo lookup user abc123
// tropo lookup user abc123 --addresses
// tropo lookup user abc123 --applications --addresses

var lookup = cli.Command{
	Name:        "lookup",
	Usage:       "tropo lookup 9fb9f0887171a133e4ce14025baa968e",
	Description: `Lookup info on token or session guid`,
	Subcommands: []cli.Command{
		{
			Name:  "application",
			Usage: "Lookup information on an application",
			Flags: []cli.Flag{
				cli.BoolFlag{"addresses, a", "Include all mapped addresses"},
				cli.BoolFlag{"pin, p", "Include mapped pins"},
				cli.BoolFlag{"tokens, t", "Include mapped tokens"},
				cli.BoolFlag{"numbers, n", "Include mapped numbers"},
			},
			Action: LookupApplication,
		},
		{
			Name:   "address",
			Usage:  "Lookup information on an address",
			Action: LookupAddress,
		},
		{
			Name:  "user",
			Usage: "Lookup a user data",
			Flags: []cli.Flag{
				cli.BoolFlag{"applications, a", "Include users applications"},
			},
			Action: LookupUser,
		},
	},
	Flags: []cli.Flag{
		cli.StringFlag{"api", "hosted", "Api to use, default `hosted`"},
	},
}

var guid = cli.Command{
	Name:        "guid",
	Usage:       "tropo guid 9fb9f0887171a133e4ce14025baa968e",
	Description: `Decode runtime IP address from session guid`,
	Action:      DecodeSessionID,
}

var list = cli.Command{
	Name:        "list",
	Usage:       "tropo list features",
	Description: ``,
	Subcommands: []cli.Command{
		{
			Name:   "features",
			Usage:  "Print feature flags",
			Action: PrintFeatures,
		},
		{
			Name:   "ppids",
			Usage:  "List all PPIDs",
			Action: PrintPpids,
		},
	},
}
