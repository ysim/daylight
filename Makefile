#!/usr/bin/make -f

# Expanded at declaration time
export GOPATH:=$(CURDIR)

# Set if absent; this is the default for OS X
export GOROOT?=/usr/local/go

all: build

clean:
	$(RM) daylight

run:
	go run main.go

test:
	go test

build: main.go
	go build -o daylight

.PHONY: all clean run test build
