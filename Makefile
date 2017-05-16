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

build: main.go
	go build -o daylight
