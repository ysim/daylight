package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var apiUrl = "https://api.sunrise-sunset.org/json"
var externalIPurl = "http://checkip.amazonaws.com/"

func checkIP() string {
	response, err := http.Get(externalIPurl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// html is uint8
	return string(bytes)
}

func main() {
	fmt.Println("api endpoint:", apiUrl)
	fmt.Printf("%s", checkIP())
}
