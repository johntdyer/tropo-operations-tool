package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/robfig/config"
	"github.com/wsxiaoys/terminal"
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

func (ac *ApplicationConfig) setConfig(section string) error {
	cfg, err := config.ReadDefault(fmt.Sprintf("%s/.tropo-api.cfg", userHomeDir()))

	if err != nil {
		printError()
		createConfig()

		// Read config again
		cfg, err = config.ReadDefault(fmt.Sprintf("%s/.tropo-api.cfg", userHomeDir()))
		if err != nil {
			printError()
		}
	}

	if cfg.HasSection(section) {
		ac.Credentials.Password = validateConfig(cfg, section, "password")
		ac.Credentials.Username = validateConfig(cfg, section, "username")
		ac.API.URL = validateConfig(cfg, section, "url")
		ac.API.InsecureSkipVerify, _ = cfg.Bool(section, "insecureSkipVerify")
		return nil
	}
	return fmt.Errorf("Unable to find config section '%s'", section)

}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func removeNewLines(str, replace string) string {
	return strings.Replace(strings.Replace(str, "\r", replace, -1), "\n", replace, -1)
}

func checkForRequiredArguments(arg, msg string) {
	if arg == "" {
		terminal.Stdout.Color("r").Print("-- ERROR -- Missing argument").Nl().Reset()
		f, _ := "  Example: \n    tropo %s", msg
		terminal.Stdout.Color("b").Print(f).Nl().Reset()
		os.Exit(0)
	}
}

func setLoggerLevel(debugMode bool) {
	if debugMode {
		logger.Level = log.DebugLevel
	}
}
