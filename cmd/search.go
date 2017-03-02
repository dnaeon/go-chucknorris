package cmd

import (
	"fmt"

	"github.com/dnaeon/go-chucknorris/api"
	"github.com/urfave/cli"
)

func NewSearchCommand() cli.Command {
	cmd := cli.Command{
		Name:   "search",
		Usage:  "search for jokes matching a given query",
		Action: execSearchCommand,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "query",
				Usage: "the search query to use",
			},
		},
	}

	return cmd
}

func execSearchCommand(c *cli.Context) error {
	if !c.IsSet("query") {
		return cli.NewExitError("missing query flag", 64)
	}

	query := c.String("query")
	client := api.NewClient(nil)
	result, err := client.Search(query)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	fmt.Printf("Found %d joke(s) matching the given query.\n", result.Total)
	for _, joke := range result.Result {
		fmt.Printf("\n%s\n", joke.Value)
	}

	return nil
}
