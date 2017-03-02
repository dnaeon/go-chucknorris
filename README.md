## go-chucknorris

[![Build Status](https://travis-ci.org/dnaeon/go-chucknorris.svg?branch=master)](https://travis-ci.org/dnaeon/go-chucknorris)
[![GoDoc](https://godoc.org/github.com/dnaeon/go-chucknorris?status.svg)](https://godoc.org/github.com/dnaeon/go-chucknorris)
[![Go Report Card](https://goreportcard.com/badge/github.com/dnaeon/go-chucknorris)](https://goreportcard.com/report/github.com/dnaeon/go-chucknorris)

`go-chucknorris` is a Go API and CLI tool for the hand curated
Chuck Norris facts at https://api.chucknorris.io/.

## Installation

Install go-chucknorris using `go get`.

```
$ go get -v github.com/dnaeon/go-chucknorris
```

Or by cloning the repo and using `make(1)`.

```
$ git clone https://github.com/dnaeon/go-chucknorris
$ cd go-chucknorris && make install
```

## Usage

Example usage of retrieving a random joke.

```
$ go-chucknorris joke
If Chuck Norris wants your opinion, he'll beat it into you.
```

Listing all joke categories.

```
$ go-chucknorris categories
```

Searching for jokes.

```
$ go-chucknorris search --query "agent smith"
Found 1 joke(s) matching the given query.

Chuck Norris once interrogated agent Smith.
```

## API

Check out the API documentation at [GoDoc](https://godoc.org/github.com/dnaeon/go-chucknorris).

## License

This project is Open Source and licensed under the [BSD License](http://opensource.org/licenses/BSD-2-Clause).
