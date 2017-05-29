#!/usr/bin/make -f

all: build

get-deps:
	go get github.com/stretchr/testify/assert
	go get github.com/kardianos/govendor # Dependency management for Heroku

clean:
	$(RM) bin/*

test:
	go test

build: daylight.go cli/main.go
	go build -o bin/daylight-cli cli/main.go

install:
	install bin/daylight-cli "${HOME}/bin/daylight-cli"

deploy-web:
	git subtree push --prefix web heroku master

deploy-web-force:
	git push heroku `git subtree split --prefix web master`:master --force

.PHONY: all get-deps clean test build deploy-web
