package main

import (
	tropo "bitbucket.org/voxeolabs/go-tropo-utils"
	"github.com/codegangsta/cli"
	"github.com/wsxiaoys/terminal"
)

func PrintFeatures(c *cli.Context) {
	BuildFeaturesTable()
}

func PrintPpids(c *cli.Context) {
	BuildPpidsTable()
}

func PrintCallBacks(c *cli.Context) {
	BuildCallBackTable()
}

func DecodeSessionID(c *cli.Context) {
	session := c.Args().First()

	CheckForRequiredArguments(session, "guid 23b1a72988afe289a11800ce553ab6b6")

	s := tropo.DecodeSessionId(session)
	terminal.Stdout.Color("y").Print("Results: ", s.IP).Nl().Reset()
}

func LookupAddress(c *cli.Context) {
	address := c.Args().First()

	CheckForRequiredArguments(address, "lookup address +14075551212")

	str, json := GetAddressData(user, password, rest_api, address)
	logger.Debug("Address Response: ", str)
	terminal.Stdout.Color("y").Print("Results").Nl().Reset()
	BuildAddressTable(json)
}

func LookupUser(c *cli.Context) {
	account := c.Args().First()

	CheckForRequiredArguments(account, "lookup user 1234")

	features := []string{}
	str, userData := GetUserData(user, password, rest_api, account)
	features = GetUserFeatures(user, password, rest_api, account)
	var _ = str
	terminal.Stdout.Color("y").Print("Results").Nl().Reset()
	BuildUserTable(userData, features)
	if c.Bool("applications") == true {
		applications := GetUsersApplications(user, password, rest_api, account)
		BuildApplicationsTable(applications)
	}
}

func doLookup(c *cli.Context) {
}

func LookupApplication(c *cli.Context) {
	application := c.Args().First()

	CheckForRequiredArguments(application, "lookup application 1234")

	str, json := GetAppData(user, password, rest_api, application)
	var _ = str
	terminal.Stdout.Color("y").Print("Results").Nl().Reset()

	BuildApplicationTable(json)
	if c.Bool("addresses") == true {
		data := GetApplicationAddresses(user, password, rest_api, application)
		BuildApplicationAddressesTable(data)
	}

}
