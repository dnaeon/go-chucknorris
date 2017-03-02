package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dnaeon/gru/version"
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
func (c *Client) Categories() ([]Category, *http.Response, error) {
	url := fmt.Sprintf("%s/jokes/categories", c.BaseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, resp, errors.New("Unable to fetch categories")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}

	var categories []Category
	if err := json.Unmarshal(body, &categories); err != nil {
		return nil, resp, err
	}

	return categories, resp, nil
}

// RandomJoke returns a random joke.
func (c *Client) RandomJoke(categories ...Category) (*Joke, *http.Response, error) {
	// If we have categories, retrieve a joke from the first category only.
	var category Category
	if len(categories) > 0 {
		category = categories[0]
	}

	url := fmt.Sprintf("%s/jokes/random", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("User-Agent", c.UserAgent)

	if category != "" {
		values := req.URL.Query()
		values.Add("category", string(category))
		req.URL.RawQuery = values.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, resp, errors.New("Unable to fetch joke")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}

	var joke Joke
	if err := json.Unmarshal(body, &joke); err != nil {
		return nil, resp, err
	}

	return &joke, resp, nil
}
