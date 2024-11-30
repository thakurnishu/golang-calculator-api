.DEFAULT_GOAL := run

build:
	@go build -o bin/main ./cmd

run: build
	@./bin/main
