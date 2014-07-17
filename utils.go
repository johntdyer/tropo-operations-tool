package main

import (
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

func GetPapiConfig() (string, string, string) {
	section := "hosted"
	var user, pass, url string

	c, err := config.ReadDefault(fmt.Sprintf("%s/.tropo-api.cfg", UserHomeDir()))
	if err != nil {
		PrintError()
		CreateConfig()
		return GetPapiConfig()
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
