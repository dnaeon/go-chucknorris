// Copyright (c) 2017 Marin Atanasov Nikolov <dnaeon@gmail.com>
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
//
//  1. Redistributions of source code must retain the above copyright
//     notice, this list of conditions and the following disclaimer
//     in this position and unchanged.
//  2. Redistributions in binary form must reproduce the above copyright
//     notice, this list of conditions and the following disclaimer in the
//     documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE AUTHOR(S) ``AS IS'' AND ANY EXPRESS OR
// IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
// OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE AUTHOR(S) BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
// THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package cmd

import (
	"fmt"

	"github.com/dnaeon/go-chucknorris/api"
	"github.com/urfave/cli"
)

// NewSearchCommand creates a new command for searching jokes.
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
