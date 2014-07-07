package main

import (
  "github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandAddressLookup,
  commandUserLookup,
  commandApplicationLookup,
  sessionIdLookup,
  list,
}

var sessionIdLookup = cli.Command{
  Name:  "guid",
  Usage: "tropo guid 9fb9f0887171a133e4ce14025baa968e",
  Description: `Decode runtime IP address from session guid`,
  Action: DecodeSessionID,
}

var list = cli.Command{
  Name:  "list",
  Usage: "tropo list features",
  Description: ``,
  Subcommands: []cli.Command{
    {
      Name:  "features",
      Usage: "Print feature flags",
      Action: PrintFeatures,
    },
  },
}
