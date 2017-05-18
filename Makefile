#!/usr/bin/make -f

# Expanded at declaration time; references Go workspace four directories up
export GOPATH:=$(CURDIR)/../../../../

# Set if absent; this is the default for OS X
export GOROOT?=/usr/local/go

all: build

get-deps:
	go get github.com/stretchr/testify/assert

clean:
	$(RM) bin/*

test:
	go test

build: daylight.go cli/main.go
	go build -o bin/daylight-cli cli/main.go
	ls -al bin/*

.PHONY: all get-deps clean run test build
