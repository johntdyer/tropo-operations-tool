package main

import (
	"github.com/Sirupsen/logrus"
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
			logger.Level = logrus.Debug
		}
	}
}
