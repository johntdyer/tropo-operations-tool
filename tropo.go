package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/op/go-logging"
	"gopkg.in/alecthomas/kingpin.v2"
)

func addSubCommand(app *kingpin.Application, name string, description string) {
	app.Command(name, description).Action(func(c *kingpin.ParseContext) error {
		fmt.Printf("Would have run command %s.\n", name)
		return nil
	})
}

var (
	// Logger
	logger = logrus.New()

	tropoAppConfig = ApplicationConfig{}
	app            = kingpin.New("tropo", "Tropo operations utility")
	debug          = app.Flag("debug", "debug logging").Bool()

	// Global Configs
	configContext = app.Flag("config", "config to use in ~/.tropo-api.cfg").Short('c').Default("hosted").String()

	// Decode GUID Commands
	cmdDecodeSessionID = app.Command("guid", "Decode SessionGUID.")
	sessionGUID        = cmdDecodeSessionID.Arg("session", "sessionId").Required().String()

	// List PPIDs
	ppids = app.Command("ppids", "List all PPIDS")

	// list feature codes
	features = app.Command("features", "List all feature codes")

	// List and lookup SIP codes
	cmdSIPCodeLookup = app.Command("sip", "Lookup sip code")
	sipCode          = cmdSIPCodeLookup.Arg("sip-code", "sip-code").String()
	includeRFCS      = cmdSIPCodeLookup.Flag("include-rfc", "include RFC's in output").Short('r').Bool()

	//
	// Lookup Stuff
	lookup = app.Command("lookup", "lookup data")

	// Application Lookup
	cmdApplication   = lookup.Command("application", "").Alias("app")
	applicationID    = cmdApplication.Arg("applicationID", "applicationID").Required().String()
	includeAddresses = cmdApplication.Flag("address", "Include addresses").Short('a').Bool()

	// Address lookup
	cmdAddress = lookup.Command("address", "").Alias("addr")
	address    = cmdAddress.Arg("address", "address to lookup").Required().String()

	// User Lookup
	cmdUser             = lookup.Command("user", "").Alias("u")
	user                = cmdUser.Arg("user", "UserID or UserName").Required().String()
	includeApplications = cmdUser.Flag("applications", "Include addresses").Short('a').Bool()
)

func init() {
	logger.Formatter = new(logrus.JSONFormatter)
	logger.Formatter = new(logrus.TextFormatter) // default
	tropoAppConfig = ApplicationConfig{}
}

// PAPIConfigPreAction - Setup PAPI Config
func PAPIConfigPreAction(c *kingpin.ParseContext) error {

	// Set log level
	setLoggerLevel(*debug)

	logger.Debug("Pre config setup of PAPI config")
	err := tropoAppConfig.setConfig(*configContext)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Debug("Pre config setup of PAPI config [ DONE ]")
	return nil
}

func main() {
	logging.SetLevel(logging.ERROR, "tropo")

	app.Version(Version).
		Author("John Dyer <johntdyer@gmail.com>").
		PreAction(PAPIConfigPreAction)

	logger.Debug("Starting application ")

	// addSubCommand(app, "ping", "Additional top level command to show command completion")
	// addSubCommand(app, "nmap", "Additional top level command to show command completion")

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case cmdDecodeSessionID.FullCommand():
		decodeSessionIDAction(*sessionGUID)
	case "lookup application":
		applicationLookupAction(*applicationID)
	case "features":
		buildFeaturesTable()
	case "lookup user":
		userLookupAction(*user)
	case "lookup address":
		addressLookupAction(*address)
	case "sip":
		sipCodeLookupAction(*sipCode)
	case "ppids":
		buildPpidsTable()
	default:
		fmt.Println("unknown acction")

	}
}
