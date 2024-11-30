.DEFAULT_GOAL := run

fmt: 
	@go fmt ./...

build: fmt
	@mkdir -p bin
	@go build -o bin/main ./cmd

run: build
	@./bin/main
