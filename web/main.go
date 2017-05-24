package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func GetClientIP(r *http.Request) (string, string) {
	return r.RemoteAddr, r.Header.Get("X-Forwarded-For")
}

func handler(w http.ResponseWriter, r *http.Request) {
	remoteAddr, originIP := GetClientIP(r)
	fmt.Fprintf(w, "RemoteAddr: %s, X-Forwarded-For: %s", remoteAddr, originIP)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
