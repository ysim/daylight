package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var apiUrl = "https://api.sunrise-sunset.org/json"
var externalIPurl = "http://checkip.amazonaws.com/"

func checkIP() {
	response, err := http.Get(externalIPurl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", html)
}

func main() {
	fmt.Println("api endpoint:", apiUrl)
	checkIP()
}
