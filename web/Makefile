.PHONY: fetch-deps
fetch-deps:
	go get github.com/kardianos/govendor # Dependency management for Heroku
	govendor fetch github.com/ysim/daylight

.PHONY: run
run:
	DAYLIGHT_ENV=local PORT=8080 go run main.go
