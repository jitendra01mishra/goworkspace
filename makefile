# Variables
BINARY := goworkspace
VERBOSE ?= false
PKG := ./...
SRC := $(shell find . -type f -name '*.go')

# Phony targets
.PHONY: all build test fmt lint clean run deps generate run-dev run-prod

# Default target
all: build

# Build the binary
build:
ifeq ($(VERBOSE),true)
	go build -v -o $(BINARY) main.go
else
	go build -o $(BINARY) main.go
endif

# Run tests
test:
	go test -v $(PKG)

# Format source code
fmt:
	go fmt $(PKG)

# Run linters (using 'staticcheck' as an example)
lint:
	staticcheck $(PKG)

deps:
	go mod tidy

generate:
	go generate $(PKG)

# Clean build artifacts
clean:
	rm -f $(BINARY)

# Run Dev the application
run-dev: build
	ENV=dev ./$(BINARY)

# Run Prod the application
run-prod: build
	ENV=prod ./$(BINARY)

GOFILES := $(wildcard *.go)
%.o: %.go
	go build -o $@ $<

# End of Makefile