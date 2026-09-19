BINARY_NAME ?= boilerplate

.PHONY: build run test lint fmt vet

build:
	go build -trimpath -o $(BINARY_NAME) .

run:
	go run .

test:
	go test -race ./...

vet:
	go vet ./...

lint: vet
	golangci-lint run

fmt:
	gofmt -w .
