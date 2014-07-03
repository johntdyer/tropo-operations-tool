package main

import (
  "time"
	"fmt"
  "net/http"
  "io/ioutil"
  "strings"
  "os"
  "code.google.com/p/gcfg"
  "log"
  "github.com/wsxiaoys/terminal"
  "github.com/olekukonko/tablewriter"
  "encoding/json"
  "strconv"
  "github.com/codegangsta/cli"
  "lib"
)



// Tries to find out when this binary was compiled.
// Returns the current time if it fails to find it.
func compileTime() time.Time {
  info, err := os.Stat(os.Args[0])
  if err != nil {
    return time.Now()
  }
  return info.ModTime()
}

func buildUserTable(papi PapiUserResponse){
  fullName := []string{papi.FirstName, papi.LastName};
  address :=  []string{papi.Address, papi.Address2, papi.State}
  notes := ""

  if papi.Notes == "" {
    notes = "none"
  }else{
    notes = papi.Notes
  }

  data := [][]string{
    []string{"Username",                papi.Username},
    []string{"AccountId",               papi.Id},
    []string{"Email",                   papi.Email},
    []string{"Name",                    strings.Join(fullName, " ")},
    []string{"Address",                 strings.Join(address, "\n")},
    []string{"JoinDate",                papi.JoinDate},
    []string{"Status",                  papi.Status},
    []string{"Notes",                   notes},
    []string{"PasswordFailedAttempts",  strconv.Itoa(papi.PasswordFailedAttempts)},
  }

  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Property", "Value"})

  for _, v := range data {
    table.Append(v)
  }
  table.SetAlignment(tablewriter.ALIGN_LEFT)

  table.Render() // Send output
}


func buildApplicationTable(papi PapiApplicationResponse){
  data := [][]string{
    []string{"AppId",         papi.Id},
    []string{"UserId",        strconv.Itoa(papi.UserId)},
    []string{"App Name",      papi.Name},
    []string{"Platform",      papi.Platform},
    []string{"Environment",   papi.Environment},
    []string{"MessagingUrl",  papi.MessagingUrl},
    []string{"VoiceUrl",      papi.VoiceUrl},
    []string{"Partition",     papi.Partition},
  }
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Property", "Value"})

  for _, v := range data {
    table.Append(v)
  }
  table.SetAlignment(tablewriter.ALIGN_LEFT)
  table.Render() // Send output
}



func printError() {
  terminal.Stdout.Color("r").Print("ERROR - Unable to locate api configuration ( ~/.tropo-api.cfg ).\n").Nl().Reset()
  fmt.Println("Please create this file in your home directory")
  terminal.Stdout.Color("b").Print("---------------------------------\n").Nl().Reset()
  fmt.Println("   ;Tropo API Configuration\n   [api]\n   url = https://api.aws.tropo.com/rest/v1/users\n   [credentials]\n   username = <username>\n   password = <password>\n")
  terminal.Stdout.Color("b").Print("---------------------------------\n").Nl().Reset()
}

func getConfig() (string, string, string){
  var cfg Config
  err := gcfg.ReadFileInto(&cfg, "/Users/jdyer/.tropo-api.gcfg")
  if err != nil {
    printError()
    os.Exit(1)
  }
  return cfg.Credentials.Username, cfg.Credentials.Password, cfg.Api.Url
}

func getApplicationContent(username, passwd, url, application string) (string, PapiApplicationResponse) {

  fullApiUrl := []string{url, "/applications/", application};
  client := &http.Client{}
  req, err := http.NewRequest("GET",strings.Join(fullApiUrl, ""), nil)
  req.SetBasicAuth(username, passwd)

  resp, err := client.Do(req)
  if err != nil{ log.Fatal(err) }

  bodyText, err := ioutil.ReadAll(resp.Body)
  if err != nil { panic(err.Error()) }

  var data PapiApplicationResponse
  err = json.Unmarshal(bodyText, &data)
  if err != nil {  panic(err.Error()) }

  str := string(bodyText)
  return str, data
}

func getUserContent(u, p, url, account string) (string, PapiUserResponse) {
  var username string = u
  var passwd string = p
  var apiUrl string = url
  var accountName string = account
  fullApiUrl := []string{apiUrl, "/users/", accountName};

  client := &http.Client{}
  req, err := http.NewRequest("GET",strings.Join(fullApiUrl, ""), nil)
  req.SetBasicAuth(username, passwd)

  resp, err := client.Do(req)
  if err != nil{ log.Fatal(err) }

  bodyText, err := ioutil.ReadAll(resp.Body)
  if err != nil { panic(err.Error()) }

  var data PapiUserResponse
  err = json.Unmarshal(bodyText, &data)
  if err != nil {  panic(err.Error()) }

  str := string(bodyText)
  return str, data
}



var Commands = []cli.Command{
	commandAddressLookup,
  commandUserLookup,
  commandApplicationLookup,
}

var commandAddressLookup = cli.Command{
  Name:  "address",
  Usage: "",
  Description: ``,
  Flags: []cli.Flag {
      cli.StringFlag{"pin, p", "", "Sip pin to lookup, eg 9995551212"},
      cli.StringFlag{"number, n", "", "Number to lookup, Must include + and country code ( +14075551212 ) "},
      cli.StringFlag{"token, t", "", "Address to lookup."},
  },
  Action: doAddressLookup,
}

var commandUserLookup = cli.Command{
  Name:  "user",
  Usage: "",
  Description: ``,
  Flags: []cli.Flag {
    cli.StringFlag{"user, u", "", "account to lookup.  Both id & username are supported"},
  },
  Action: doUserLookup,
}

var commandApplicationLookup = cli.Command{
	Name:  "application",
	Usage: "",
	Description: ``,
  Flags: []cli.Flag {
    cli.StringFlag{"app, a", "", "Application ID to lookup."},
  },
	Action: doApplicationLookup,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doAddressLookup(c *cli.Context) {
  user, password, url := getConfig()

  app := cli.NewApp()
  app.Name = "tlookup"
  app.Usage = "Tropo infomatation lookup tool"

  if c.String("user") != "" {
    str, json := getUserContent(user, password, url, c.String("user"))
    var _ = str
    terminal.Stdout.Color("y").Print("Results").Nl().Reset()
    buildTable(json)
    os.Exit(0)
  }
}

func doUserLookup(c *cli.Context) {
  user, password, url := getConfig()

  if c.String("user") == "" {
    terminal.Stdout.Color("r").Print("-- ERROR -- Missing argument").Nl().Reset()
    terminal.Stdout.Color("b").Print("  Examples: \n").Nl().Reset()
    terminal.Stdout.Color("b").Print("    tlookup address --pin 9995551212").Nl().Reset()
    terminal.Stdout.Color("b").Print("    tlookup address --token cf77f59c97b9a44c950446eff97a37bc100212a34ed68a4b0115dfc79d9f3b73446b9995c9170bdec6e1a889").Nl().Reset()
    terminal.Stdout.Color("b").Print("    tlookup address --number +14075551212").Nl().Reset()
    fmt.Println("---------------")
    cli.ShowAppHelp(c)
  }else {
    str, json := getAddressContent(user, password, url, c.String("address"))
    var _ = str
    terminal.Stdout.Color("y").Print("Results").Nl().Reset()
    buildAddressTable(json)
  }
  os.Exit(0)
}

func doUserLookup(c *cli.Context) {
  user, password, url := getConfig()

  if c.String("user") == "" {
    terminal.Stdout.Color("r").Print("-- ERROR -- Missing argument").Nl().Reset()
    terminal.Stdout.Color("b").Print("  Example: \n    tlookup user --user 1234").Nl().Reset()
    fmt.Println("---------------")
    cli.ShowAppHelp(c)
  }else {
    str, json := getUserContent(user, password, url, c.String("user"))
    var _ = str
    terminal.Stdout.Color("y").Print("Results").Nl().Reset()
    buildUserTable(json)
  }
  os.Exit(0)
}

func doApplicationLookup(c *cli.Context) {
  user, password, url := getConfig()

  if c.String("app") == "" {
    terminal.Stdout.Color("r").Print("-- ERROR -- Missing argument").Nl().Reset()
    terminal.Stdout.Color("b").Print("  Example: \n    tlookup user --user 1234").Nl().Reset()
    fmt.Println("---------------")
    cli.ShowAppHelp(c)
  }else {
    str, json := getApplicationContent(user, password, url, c.String("app"))
    var _ = str
    terminal.Stdout.Color("y").Print("Results").Nl().Reset()
    buildApplicationTable(json)
  }
  os.Exit(0)

}

