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

// tot lookup address 4075551212 --user --application
// tot lookup address 99999999   --user --application
// tot lookup address abc123     --user --application
// tot lookup appplication 1234  --addresses

// tot guid abc123
// tot check dns 4075551212

// tot lookup user abc123
// tot lookup user abc123 --addresses
// tot lookup user abc123 --applications --addresses

var lookup = cli.Command{
	Name:        "lookup",
	Usage:       "tot lookup 9fb9f0887171a133e4ce14025baa968e",
	Description: `Lookup info on token or session guid`,
	Subcommands: []cli.Command{
		{
			Name:  "application",
			Usage: "Lookup information on an application",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "addresses, a", Usage: "Include all mapped addresses"},
				cli.BoolFlag{Name: "pin, p", Usage: "Include mapped pins"},
				cli.BoolFlag{Name: "tokens, t", Usage: "Include mapped tokens"},
				cli.BoolFlag{Name: "numbers, n", Usage: "Include mapped numbers"},
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
				cli.BoolFlag{Name: "applications, a", Usage: "Include users applications"},
			},
			Action: LookupUser,
		},
		{
			Name:  "sip",
			Usage: "Lookup sip code",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all, a", Usage: "List all codes"},
			},
			Action: LookupSipCode,
		},
	},
	Flags: []cli.Flag{
		cli.StringFlag{Name: "api", Value: "hosted", Usage: "Api to use, default `hosted`"},
	},
}

var guid = cli.Command{
	Name:        "guid",
	Usage:       "tot guid 9fb9f0887171a133e4ce14025baa968e",
	Description: `Decode runtime IP address from session guid`,
	Action:      DecodeSessionID,
}

var list = cli.Command{
	Name:        "list",
	Usage:       "tot list features",
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
		{
			Name:   "callbacks",
			Usage:  "List sip codes and their tropo callbacks",
			Action: PrintAllCallBacks,
		},
	},
}
