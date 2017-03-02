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

package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dnaeon/go-chucknorris/version"
)

const (
	userAgent      = "go-chucknorris/" + version.Version
	defaultBaseUrl = "https://api.chucknorris.io"
)

// Category represents a joke category.
type Category string

// Joke type represents a Chuck Norris joke.
type Joke struct {
	// IconURL is the url to the Chuck Norris API icon.
	IconURL string `json:"icon_url"`

	// Id is the id of the joke.
	Id string `json:"id"`

	// URL is the url to the joke.
	URL string `json:"url"`

	// Value contains the joke text.
	Value string `json:"value"`
}

// SearchResponse is the response returned by the API when searching for jokes.
type SearchResponse struct {
	// Total is the total number of jokes matching the search query.
	Total int `json:"total"`

	// Result is the list of jokes matching the search query.
	Result []Joke `json:"result"`
}

// Client is an API client to the Chuck Norris jokes.
type Client struct {
	// Client to use when interacting with the API.
	client *http.Client

	// BaseURL is the url to the Chuck Norris jokes API.
	BaseURL string

	// UserAgent is the user agent to use when interacting with the API.
	UserAgent string
}

// NewClient creates a new API client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		client:    httpClient,
		BaseURL:   defaultBaseUrl,
		UserAgent: userAgent,
	}

	return client
}

// Categories returns the list of categories.
func (c *Client) Categories() ([]Category, error) {
	url := fmt.Sprintf("%s/jokes/categories", c.BaseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Unable to fetch categories")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var categories []Category
	if err := json.Unmarshal(body, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

// RandomJoke returns a random joke.
func (c *Client) RandomJoke(categories ...Category) (*Joke, error) {
	// If we have categories, retrieve a joke from the first category only.
	var category Category
	if len(categories) > 0 {
		category = categories[0]
	}

	url := fmt.Sprintf("%s/jokes/random", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.UserAgent)

	if category != "" {
		values := req.URL.Query()
		values.Add("category", string(category))
		req.URL.RawQuery = values.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Unable to fetch joke")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var joke Joke
	if err := json.Unmarshal(body, &joke); err != nil {
		return nil, err
	}

	return &joke, nil
}

// Search searches for jokes matching a given search query.
func (c *Client) Search(query string) (*SearchResponse, error) {
	url := fmt.Sprintf("%s/jokes/search", c.BaseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.UserAgent)

	values := req.URL.Query()
	values.Add("query", query)
	req.URL.RawQuery = values.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Unable to search for jokes")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result SearchResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
