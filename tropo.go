package main

import (
	"github.com/codegangsta/cli"
	"github.com/op/go-logging"
	"os"
)

var user, password, rest_api = GetPapiConfig()

var logger = logging.MustGetLogger("tropo")

func main() {
	logging.SetLevel(logging.ERROR,"tropo")

	app := cli.NewApp()
	app.Name     = "tropo"
	app.Version  = Version
	app.Usage    = "kicking ass and taking names"
	app.Author   = "John Dyer"
	app.Email    = "johntdyer@gmail.com"
	app.Compiled = CompileTime()
	app.Commands = Commands

	SetLoggerLevel(os.Getenv("DEBUG_MODE"))

  logger.Debug("Starting application ")

	app.Run(os.Args)
}
