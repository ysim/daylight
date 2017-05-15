package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type GeoIP struct {
	Ip          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name""`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	Zipcode     string  `json:"zipcode"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
	AreaCode    int     `json:"area_code"`
}

var geo GeoIP

var externalIPurl = "http://checkip.amazonaws.com/"
var geoIPurl = "https://freegeoip.net/json/"
var apiUrl = "https://api.sunrise-sunset.org/json"

func GetIP() string {
	response, err := http.Get(externalIPurl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(bytes))
}

func GetCoordinatesFromIP(ip string) (float64, float64) {
	response, err := http.Get(geoIPurl + ip)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(response.Body).Decode(&geo)
	latitude := geo.Latitude
	longitude := geo.Longitude
	return latitude, longitude
}

func main() {
	ipAddress := GetIP()
	latitude, longitude := GetCoordinatesFromIP(ipAddress)
	fmt.Println(latitude, longitude)
}
