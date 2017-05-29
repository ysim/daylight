.PHONY: fetch-deps
fetch-deps:
	govendor fetch github.com/ysim/daylight

.PHONY: run
run:
	DAYLIGHT_ENV=local PORT=8080 go run main.go
