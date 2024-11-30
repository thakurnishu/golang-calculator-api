.DEFAULT_GOAL := run

build:
	@mkdir -p bin
	@go build -o bin/main ./cmd

run: build
	@./bin/main
