.PHONY: build build-run test run lint

build:
	go build -o ./build/parser ./cmd/parser/main.go

test:
	go test -v -coverprofile cover.out ./... && go tool cover -html=cover.out

lint:
	golangci-lint run