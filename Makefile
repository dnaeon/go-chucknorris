build: get
	go build -v ./...

get:
	go get -v -t -d ./...

test:
	go test -v ./...

install:
	go install -v ./...

format:
	go fmt .

.PHONY: build get test install format
