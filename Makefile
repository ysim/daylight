#!/usr/bin/make -f

# Expanded at declaration time
export GOPATH:=$(CURDIR)

# Set if absent; this is the default for OS X
export GOROOT?=/usr/local/go

all: build

get-deps:
	go get github.com/stretchr/testify/assert

clean:
	$(RM) daylight

run:
	go run main.go

test:
	go test

build: main.go
	go build -o daylight

.PHONY: all get-deps clean run test build
