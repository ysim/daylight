#!/usr/bin/make -f

.PHONY: all
all: build

.PHONY: get-deps
get-deps:
	go get github.com/stretchr/testify/assert
	go get github.com/kardianos/govendor # Dependency management for Heroku

.PHONY: clean
clean:
	$(RM) bin/*

.PHONY: test
test:
	go test

.PHONY: build
build: daylight.go cli/main.go
	go build -o bin/daylight-cli cli/main.go

.PHONY: install
install:
	install bin/daylight-cli "${HOME}/bin/daylight-cli"

.PHONY: deploy-web
deploy-web:
	git subtree push --prefix web heroku master

.PHONY: deploy-web-force
deploy-web-force:
	git push heroku `git subtree split --prefix web master`:master --force
