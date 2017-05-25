package main

import (
	"fmt"
	"github.com/ysim/daylight"
	"log"
	"net/http"
	"os"
)

var daylightEnv string

func GetClientIP(r *http.Request) string {
	if daylightEnv == "local" {
		return daylight.GetIP()
	}
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		return xff
	}
	return r.RemoteAddr
}

func handler(w http.ResponseWriter, r *http.Request) {
	clientIP := GetClientIP(r)
	location := daylight.BuildLocation("", clientIP)
	location.GetSunriseSunset("today")
	location.GetLocalizedSunriseSunset()
	fmt.Fprintf(w, "%s", location.GetDisplayString())
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	daylightEnv = os.Getenv("DAYLIGHT_ENV")
	if daylightEnv == "" {
		log.Fatal("$DAYLIGHT_ENV must be set")
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
