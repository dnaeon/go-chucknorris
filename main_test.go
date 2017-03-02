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

package main

import (
	"net/http"
	"testing"

	"github.com/dnaeon/go-chucknorris/api"
	"github.com/dnaeon/go-vcr/recorder"
)

func TestMain(t *testing.T) {
	r, err := recorder.New("fixtures/chucknorris")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()

	wantJoke := &api.Joke{
		IconURL: "https://assets.chucknorris.host/img/avatar/chuck-norris.png",
		Id:      "nj19y6OkQSu7n0tYiPn4FA",
		URL:     "http://api.chucknorris.io/jokes/nj19y6OkQSu7n0tYiPn4FA",
		Value:   "Chuck Norris once interrogated agent Smith.",
	}

	httpClient := &http.Client{
		Transport: r,
	}
	client := api.NewClient(httpClient)
	result, err := client.Search("Chuck Norris once interrogated agent Smith.")
	if err != nil {
		t.Fatal(err)
	}

	if result.Total != 1 {
		t.Fatalf("want 1 joke, got %d joke(s)", result.Total)
	}

	gotJoke := result.Result[0]

	if wantJoke.IconURL != gotJoke.IconURL {
		t.Fatalf("want icon url %q, got %q", wantJoke.IconURL, gotJoke.IconURL)
	}

	if wantJoke.Id != gotJoke.Id {
		t.Fatalf("want id %q, got %q", wantJoke.Id, gotJoke.Id)
	}

	if wantJoke.URL != gotJoke.URL {
		t.Fatalf("want url %q, got %q", wantJoke.URL, gotJoke.URL)
	}

	if wantJoke.Value != gotJoke.Value {
		t.Fatalf("want value %q, got %q", wantJoke.Value, gotJoke.Value)
	}
}
