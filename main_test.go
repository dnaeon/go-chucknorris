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
