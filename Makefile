.PHONY: build run test lint lint-fix

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

run:
	go run ./cmd/hexlet-path-size $(ARGS)

test:
	go test ./...

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix
