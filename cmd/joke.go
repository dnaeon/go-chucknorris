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

// NewJokeCommand creates a new command for retrieving a Chuck Norris joke.
func NewJokeCommand() cli.Command {
	cmd := cli.Command{
		Name:   "joke",
		Usage:  "get a random joke or from given category",
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
