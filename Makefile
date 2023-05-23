BINARY_NAME=boilerplate

build:
	go build -o ${BINARY_NAME}

run: build
	./${BINARY_NAME}

test:
	go vet
	go test ./...

lint:
	golangci-lint run
