.PHONY: build run

build:
	mkdir -p bin/linux bin/darwin
	go fmt ./...
	go get ./...
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -o bin/linux/tablecheck-pinger ./cmd
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/tablecheck-pinger ./cmd
