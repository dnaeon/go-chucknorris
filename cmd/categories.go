package cmd

import (
	"fmt"

	"github.com/dnaeon/go-chucknorris/api"
	"github.com/urfave/cli"
)

// NewCategoriesCmd creates a new command for retrieving the
// Chuck Norris jokes category list.
func NewCategoriesCommand() cli.Command {
	cmd := cli.Command{
		Name:   "categories",
		Usage:  "shows the list of joke categories",
		Action: execCategoriesCommand,
	}

	return cmd
}

func execCategoriesCommand(c *cli.Context) error {
	client := api.NewClient(nil)
	categories, err := client.Categories()

	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	for _, category := range categories {
		fmt.Println(category)
	}

	return nil
}
