VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")

## install: Install all dependencies
install:
	go install

## test: Run all unit-tests
test:
	@echo "  >  run unit-tests..."
	go test -v -race -timeout 30s ./...

## bench: Run all benchmark-tests
bench:
	@echo "  >  run benchmarks..."
	go test -benchmem -bench=. ./...

## cover: Check test coverage
cover:
	@echo "  >  check coverage..."
	go test -v -cover ./...

all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: all clean install uninstall test build help gmail-connection
.DEFAULT_GOAL := help


