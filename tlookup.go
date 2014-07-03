package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name     = "tlookup"
	app.Version  = Version
	app.Usage    = ""
	app.Author   = "John Dyer"
	app.Email    = "johntdyer@gmail.com"
	app.Compiled =     CompileTime()
	app.Commands = Commands

	app.Run(os.Args)
}
