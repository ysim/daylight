package main

import (
	"fmt"
	"github.com/ysim/daylight"
	"log"
	"net/http"
	"os"
)

func GetClientIP(r *http.Request) string {
	// If request is relative (localhost), check external IP instead
	if r.URL.IsAbs() == false {
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

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
