package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/op/go-logging"
	"os"
)

var user, password, rest_api string
var InsecureSkipVerify = false

var logger = logging.MustGetLogger("tropo")

func main() {
	logging.SetLevel(logging.ERROR, "tropo")

	app := cli.NewApp()
	app.Name = "tropo"
	app.Version = Version
	app.Usage = "Tropo operations utility"
	app.Author = "John Dyer"
	app.Email = "johntdyer@gmail.com"
	app.Compiled = CompileTime()
	app.Commands = Commands
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "config, c", Value: "hosted", Usage: "config to use in ~/.tropo-api.cfg"},
	}

	app.Before = func(c *cli.Context) error {
		str := fmt.Sprintf("%s", c.String("config"))
		user, password, rest_api, InsecureSkipVerify = GetPapiConfig(str)
		return nil
	}

	SetLoggerLevel(os.Getenv("DEBUG_MODE"))

	logger.Debug("Starting application ")

	app.Run(os.Args)
}
