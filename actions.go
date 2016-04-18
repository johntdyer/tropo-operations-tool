package main

import (
	"bitbucket.org/voxeolabs/go-tropo-utils"
	"github.com/codegangsta/cli"
	"github.com/wsxiaoys/terminal"
)

// addressLookupAction
func addressLookupAction(address string) {
	str, addressData := getAddressData(address)
	logger.Debug("Address Response: ", str)
	terminal.Stdout.Color("y").Print("Results").Nl().Reset()
	buildAddressTable2(addressData)

	// if addressData.ApplicationID != 0 {
	// 	_, appData := getAppData(strconv.Itoa(addressData.ApplicationID))
	// 	features := getUserFeatures(strconv.Itoa(appData.UserID))
	// 	_, userData := getUserData(strconv.Itoa(appData.UserID))
	//
	// 	buildApplicationTable(appData)
	// 	buildUserTable(userData, features)
	// }
}

// sipCodeLookupAction
func sipCodeLookupAction(code string) {
	var showAllSIPCodes bool
	if *sipCode == "" {
		showAllSIPCodes = true
	} else {
		showAllSIPCodes = false
	}
	buildSipTable(*sipCode, showAllSIPCodes, *includeRFCS)
}

// decodeSessionIDAction - Decode a sessionGUID
func decodeSessionIDAction(sessionID string) {

	s := tropo.DecodeSessionID(sessionID)
	terminal.Stdout.Color("y").Print(sessionID + " -> ").Reset()
	terminal.Stdout.Color("r").Print(s.IP).Nl().Reset()
}

// userLookupAction
func userLookupAction(user string) {
	features := []string{}
	_, userData := getUserData(user)
	features = getUserFeatures(user)

	terminal.Stdout.Color("y").Print("Results").Nl().Reset()
	buildUserTable(userData, features)

	if *includeApplications {
		applications := getUsersApplications(user)
		buildApplicationsTable(applications)
	}
}

// applicationLookupAction
func applicationLookupAction(applicationID string) {
	_, json := getAppData(applicationID)
	terminal.Stdout.Color("y").Print("Results").Nl().Reset()

	buildApplicationTable(json)
	if *includeAddresses == true {
		data := getApplicationAddresses(applicationID)
		buildApplicationAddressesTable(data)
	}

}

func manageFeature(c *cli.Context) {
	if c.Args().Present() {
		userName := c.Args().First()
		if c.Bool("list") {
			buildFeaturesTable()
		}

		if true { //.Bool("show") {
			// fmt.Println(restAPI)
			getUserFeatures(userName)

		}
	}
}

// 	// if c.Bool("disable") {
//
// 	// 	str, err := provisioningApiPost(user, password, "https://api.tropo.com/v1/users/"+userName+"/features", []byte(`{"feature":"10","featureFlag":"x"}`))
//
// 	// 	if err != nil {
// 	// 		fmt.Println(err)
// 	// 	}
// 	// 	fmt.Println(str)
// 	// }
//
// 	// if c.Bool("enable") {
// 	// 	fmt.Println("Enabling")
// 	// 	str, err := provisioningApiDelete(user, password, "https://api.tropo.com/v1/users/"+userName+"/features/10")
// 	// 	if err != nil {
// 	// 		fmt.Println(err)
// 	// 	}
// 	// 	fmt.Println("Enabled")
// 	// 	fmt.Println(str)
//
// 	// }
// 	//  map[string]int{
// 	// 	"c": 1,
// 	// 	"u": 2,
// 	// 	"i": 3,
// 	// 	"s": 4,
// 	// 	"r": 5,
// 	// 	"b": 6,
// 	//  "t": 7,
// 	// 	"d": 8,
// 	// 	"w": 9,
// 	//  "x" : 10,
// 	// }
// 	fmt.Println(c.StringSlice("add"))
// 	fmt.Println(c.StringSlice("remove"))
//
// }
