#!/usr/bin/make -f

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

deploy-web:
	git subtree push --prefix web heroku master

.PHONY: all get-deps clean test build deploy-web
