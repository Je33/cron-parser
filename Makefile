.PHONY: build build-run test run lint

build:
	go build -o ./build/parser ./cmd/parser/main.go

build-run: build
	./build/parser

test:
	go test -v -coverprofile cover.out ./... && go tool cover -html=cover.out

run:
	go run -race ./cmd/parser/main.go

lint:
	golangci-lint run