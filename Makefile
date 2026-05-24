APP_NAME=lab1-tooling

.PHONY: fmt lint test build verify all

fmt:
	go fmt ./...

lint:
	golangci-lint run

build:
	go build -o bin/$(APP_NAME).exe ./cmd/app

verify:
	go mod verify

test:
	go test -race -v ./...

all: fmt verify lint test build