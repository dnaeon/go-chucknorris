package cmd

import (
	"fmt"

	"github.com/dnaeon/go-chucknorris/api"
	"github.com/urfave/cli"
)

// NewJokeCommand creates a new command for retrieving a Chuck Norris joke.
func NewJokeCommand() cli.Command {
	cmd := cli.Command{
		Name:   "joke",
		Usage:  "tells a Chuck Norris joke",
		Action: execJokeCommand,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "category",
				Usage: "choose a joke from given category",
			},
		},
	}

	return cmd
}

func execJokeCommand(c *cli.Context) error {
	client := api.NewClient(nil)
	category := api.Category(c.String("category"))
	joke, err := client.RandomJoke(category)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	fmt.Println(joke.Value)

	return nil
}
