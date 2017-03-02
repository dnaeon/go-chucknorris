package main

import (
	"os"

	"github.com/dnaeon/go-chucknorris/cmd"
	"github.com/dnaeon/go-chucknorris/version"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = version.Version
	app.EnableBashCompletion = true
	app.Usage = "The Chuck Norris jokes cli tool"
	app.Commands = []cli.Command{
		cmd.NewCategoriesCommand(),
		cmd.NewJokeCommand(),
		cmd.NewSearchCommand(),
	}

	app.Run(os.Args)
}
