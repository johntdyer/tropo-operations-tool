package main

import (
	//"code.google.com/p/gcfg"
	"fmt"
	"github.com/op/go-logging"
	"github.com/robfig/config"
	"github.com/wsxiaoys/terminal"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
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

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// type argError struct {
//     arg  int
//     prob string
// }

// func ValidateConfig(cfg Config) (int, error){
//   if cfg.Credentials.Username == ""{
//     return 1, &argError{1, "Can't display record"}
//   }
//   //if cfg.Credentials.Password == nil
//   //if cfg.Api.Url == nil
//   return 1, nil
// }

func printError() {
	home := fmt.Sprintf("ERROR - Unable to locate api configuration ( %s/.tropo-api.cfg ).\n", UserHomeDir())
	terminal.Stdout.Color("r").Print(home).Nl().Reset()
	fmt.Println("Please create this file in your home directory")
	terminal.Stdout.Color("b").Print("---------------------------------\n").Nl().Reset()
	//fmt.Println("   ;Tropo API Configuration\n   [api]\n   url = https://api.aws.tropo.com/rest/v1/users\n   [credentials]\n   username = <username>\n   password = <password>\n")
	fmt.Println("   ; Default api-config\n   ; https://github.com/robfig/config\n   [DEFAULT]\n   host: api.aws.tropo.com\n   route: /rest/v1\n   protocol: https://\n   base-url: %(protocol)s%(host)s%(route)s\n   \n   [hosted]\n   url: %(base-url)s\n   username: <username>\n   password: <password>\n")
	terminal.Stdout.Color("b").Print("---------------------------------\n").Nl().Reset()
}

// func GetPapiConfig() (string, string, string) {
// 	var cfg Config
// 	err := gcfg.ReadFileInto(&cfg, "/Users/jdyer/.tropo-api.gcfg")
// 	if err != nil {
// 		printError()
// 		os.Exit(1)
// 	}
// 	//ValidateConfig(cfg)

// 	return cfg.Credentials.Username, cfg.Credentials.Password, cfg.Api.Url
// }

func GetPapiConfig() (string, string, string) {
	section := "hosted"
	var user, pass, url string

	c, err := config.ReadDefault(fmt.Sprintf("%s/.tropo-api.cfg", UserHomeDir()))
	if err != nil {
		printError()
		os.Exit(1)
	}

	if c.HasSection(section) {

		url, err := c.String(section, "url")
		if err != nil {
			logger.Fatal(err)
			os.Exit(1)
		}

		if url == "" {
			terminal.Stdout.Color("r").Print("-- ERROR -- URL is required but was not found in config").Nl().Reset()
			os.Exit(1)
		}

		user, _ := c.String(section, "username")
		if err != nil {
			logger.Fatal(err)
			os.Exit(1)
		}

		if user == "" {
			terminal.Stdout.Color("r").Print("-- ERROR -- Username is required but was not found in config").Nl().Reset()
			os.Exit(1)
		}

		pass, _ := c.String(section, "password")
		if err != nil {
			logger.Fatal(err)
			os.Exit(1)
		}

		if pass == "" {
			terminal.Stdout.Color("r").Print("-- ERROR -- Password is required but was not found in config").Nl().Reset()
			os.Exit(1)
		}

		return user, pass, url
	} else {
		logger.Fatal("Unable to find section [ ", section, " ]")
		os.Exit(1)
	}
	return user, pass, url
}

func RemoveNewLines(str, replace string) string {
	return strings.Replace(strings.Replace(str, "\r", replace, -1), "\n", replace, -1)
}
func CheckForRequiredArguments(arg, msg string) {
	if arg == "" {
		terminal.Stdout.Color("r").Print("-- ERROR -- Missing argument").Nl().Reset()
		terminal.Stdout.Color("b").Print("  Example: \n    tropo %s", msg).Nl().Reset()
		os.Exit(0)
	}
}

func SetLoggerLevel(raw_debug string) {
	// Debug mode
	if raw_debug != "" {
		debug_mode, err := strconv.ParseBool(raw_debug)
		if err != nil {
			logger.Fatal(err)
		}

		if debug_mode {
			logging.SetLevel(logging.DEBUG, "tropo")
		}
	}
}
