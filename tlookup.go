package main

import (
	"github.com/codegangsta/cli"
	"os"
	"fmt"
	"strconv"
	"github.com/op/go-logging"
)
var LOGGER = logging.MustGetLogger("tlookup")

func main() {

	app := cli.NewApp()
	app.Name     = "tlookup"
	app.Version  = Version
	app.Usage    = "kicking ass and taking names"
	app.Author   = "John Dyer"
	app.Email    = "johntdyer@gmail.com"
	app.Compiled = CompileTime()
	app.Commands = Commands

	// app.Flags = []cli.Flag {
 //    cli.BoolFlag{"debug, d", "Run in debug mode"},
 //  }

	// app.Action = func(c *cli.Context) {
	// 	foo := strconv.FormatBool(c.GlobalBool("debug"))
	// 	fmt.Println("====> ",  foo) // TNothing ever happens here
 //  }

	app.Run(os.Args)
}
