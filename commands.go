package main

import (
	"fmt"
  "github.com/codegangsta/cli"
  "github.com/wsxiaoys/terminal"
  tropo "bitbucket.org/voxeolabs/go-tropo-utils"
)

var Commands = []cli.Command{
	commandAddressLookup,
  commandUserLookup,
  commandApplicationLookup,
  sessionIdLookup,
}

var sessionIdLookup = cli.Command{
  Name:  "guid",
  Usage: "tropo guid 9fb9f0887171a133e4ce14025baa968e",
  Description: `Decode runtime IP address from session guid`,
  Action: doSessionLookup,
}

var commandAddressLookup = cli.Command{
  Name:  "address",
  Usage: "tropo address --number +14074740214",
  Description: `Lookup adddress data`,
  Flags: []cli.Flag {
    cli.StringFlag{"pin, p", "", "Sip pin to lookup, eg 9995551212"},
    cli.StringFlag{"number, n", "", "Number to lookup, Must include + and country code ( +14075551212 ) "},
    cli.StringFlag{"token, t", "", "Address to lookup."},
  },
  Action: doAddressLookup,
}

var commandUserLookup = cli.Command{
  Name:  "user",
  Usage: "tropo user --user jdyer",
  Description: ``,
  Flags: []cli.Flag {
    cli.StringFlag{"user, u", "", "account to lookup.  Both id & username are supported"},
  },
  Action: doUserLookup,
}

var commandApplicationLookup = cli.Command{
	Name:  "application",
	Usage: "tropo application --app 123456",
	Description: `Lookup application data `,
  Flags: []cli.Flag {
    cli.StringFlag{"app, a", "", "Application ID to lookup."},
  },
	Action: doApplicationLookup,
}

func assert(err error) {
	if err != nil {
		logger.Fatal(err)
	}
}

func doSessionLookup(c *cli.Context) {
  s := tropo.DecodeSessionId(c.Args()[0])
  terminal.Stdout.Color("y").Print("Results: ", s.IP).Nl().Reset()
}

func doAddressLookup(c *cli.Context) {
  user, password, url := GetPapiConfig()

  if c.String("ff") != "" {
  //   terminal.Stdout.Color("r").Print("-- ERROR -- Missing argument").Nl().Reset()
  //   terminal.Stdout.Color("b").Print("  Examples: \n").Nl().Reset()
  // else if c.String("number") == "" {
  //   terminal.Stdout.Color("r").Print("-- ERROR -- Missing argument").Nl().Reset()
  //   terminal.Stdout.Color("b").Print("  Examples: \n").Nl().Reset()
  //   terminal.Stdout.Color("b").Print("    tropo address --number +14075551212").Nl().Reset()
  //   }


  //   terminal.Stdout.Color("b").Print("    tropo address --pin 9995551212").Nl().Reset()
  //   terminal.Stdout.Color("b").Print("    tropo address --token cf77f59c97b9a44c950446eff97a37bc100212a34ed68a4b0115dfc79d9f3b73446b9995c9170bdec6e1a889").Nl().Reset()

  //   fmt.Println("---------------")
  //   cli.ShowAppHelp(c)
  }else {
    str, json := GetAddressData(user, password, url, c.String("number"))
    logger.Debug("Address Response: ", str)
    terminal.Stdout.Color("y").Print("Results").Nl().Reset()
    BuildAddressTable(json)
  }
}

func doUserLookup(c *cli.Context) {
  user, password, url := GetPapiConfig()

  if c.String("user") == "" {
    terminal.Stdout.Color("r").Print("-- ERROR -- Missing argument").Nl().Reset()
    terminal.Stdout.Color("b").Print("  Example: \n    tropo user --user 1234").Nl().Reset()
    fmt.Println("---------------")
    cli.ShowAppHelp(c)
  }else {
    str, json := GetUserData(user, password, url, c.String("user"))
    var _ = str
    terminal.Stdout.Color("y").Print("Results").Nl().Reset()
    BuildUserTable(json)
  }
}

func doApplicationLookup(c *cli.Context) {
  user, password, url := GetPapiConfig()

  if c.String("app") == "" {
    terminal.Stdout.Color("r").Print("-- ERROR -- Missing argument").Nl().Reset()
    terminal.Stdout.Color("b").Print("  Example: \n    tropo user --user 1234").Nl().Reset()
    fmt.Println("---------------")
    cli.ShowAppHelp(c)
  }else {
    str, json := GetAppData(user, password, url, c.String("app"))
    var _ = str
    terminal.Stdout.Color("y").Print("Results").Nl().Reset()
    BuildApplicationTable(json)
  }

}

