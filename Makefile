.PHONY: help web test lint

help: # me
	@grep '^[-a-z.]*:' Makefile | sed -e 's/^\(.*\): .*# \(.*\)/\1 - \2/'

web: # Run web server.
	air

test: # Run all tests.
	go test -v ./internal/*

lint: # run linter.
	golangci-lint run ./cmd/web
