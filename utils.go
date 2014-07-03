package main

import(
  "fmt"
  "os"
  "time"
  "code.google.com/p/gcfg"
  "github.com/wsxiaoys/terminal"
)

// Tries to find out when this binary was compiled.
// Returns the current time if it fails to find it.
func CompileTime() time.Time {
  info, err := os.Stat(os.Args[0])
  if err != nil {
    return time.Now()
  }
  return info.ModTime()
}

func printError() {
  terminal.Stdout.Color("r").Print("ERROR - Unable to locate api configuration ( ~/.tropo-api.cfg ).\n").Nl().Reset()
  fmt.Println("Please create this file in your home directory")
  terminal.Stdout.Color("b").Print("---------------------------------\n").Nl().Reset()
  fmt.Println("   ;Tropo API Configuration\n   [api]\n   url = https://api.aws.tropo.com/rest/v1/users\n   [credentials]\n   username = <username>\n   password = <password>\n")
  terminal.Stdout.Color("b").Print("---------------------------------\n").Nl().Reset()
}

func GetPapiConfig() (string, string, string){
  var cfg Config
  err := gcfg.ReadFileInto(&cfg, "/Users/jdyer/.tropo-api.gcfg")
  if err != nil {
    printError()
    os.Exit(1)
  }
  return cfg.Credentials.Username, cfg.Credentials.Password, cfg.Api.Url
}
