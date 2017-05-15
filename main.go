package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

var apiUrl = "https://api.sunrise-sunset.org/json"
var externalIPurl = "http://checkip.amazonaws.com/"

func GetIP() net.IP {
	response, err := http.Get(externalIPurl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return net.ParseIP(strings.TrimSpace(string(bytes)))
}

func main() {
	ipAddress := GetIP()
	fmt.Println(ipAddress)
}
