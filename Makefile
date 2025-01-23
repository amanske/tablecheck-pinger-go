.PHONY: build run

build:
	mkdir -p bin
	go fmt ./...
	go get ./...
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -o bin/tablecheck-pinger-linux ./cmd
	GOOS=darwin GOARCH=amd64 go build -o bin/tablecheck-pinger-darwin ./cmd
