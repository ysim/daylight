package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	externalIPurl    = "http://checkip.amazonaws.com/"
	geoIPurl         = "https://freegeoip.net/json/"
	sunriseSunsetUrl = "https://api.sunrise-sunset.org/json"
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
	Timezone    string  `json:"time_zone"`
}

var geo GeoIP

type SunriseSunsetResults struct {
	Sunrise               string `json:"sunrise"`
	Sunset                string `json:"sunset"`
	SolorNoon             string `json:"solar_noon"`
	DayLength             string `json:"day_length"`
	CivilTwilightBegin    string `json:"civil_twilight_begin"`
	CivilTwilightEnd      string `json:"civil_twilight_end"`
	NauticalTwilightBegin string `json:"nautical_twilight_begin"`
	NauticalTwilightEnd   string `json:"nautical_twilight_end"`
	AstroTwilightBegin    string `json:"astronomical_twilight_begin"`
	AstroTwilightEnd      string `json:"astronomical_twilight_end"`
}

type SunriseSunset struct {
	Results SunriseSunsetResults `json:"results"`
	Status  string               `json:"status"`
}

var ss SunriseSunset

type Location struct {
	City      string
	Country   string
	Latitude  float64
	Longitude float64
	Timezone  string
	Sunrise   string
	Sunset    string
	DayLength string
}

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

func GetCoordinatesFromIP(ip string) Location {
	response, err := http.Get(geoIPurl + ip)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(response.Body).Decode(&geo)
	location := Location{City: geo.City, Country: geo.CountryName, Timezone: geo.Timezone, Latitude: geo.Latitude, Longitude: geo.Longitude}
	return location
}

func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', 7, 64)
}

func (location *Location) GetSunriseSunset() {
	// Build the request
	client := &http.Client{}
	req, err := http.NewRequest("GET", sunriseSunsetUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("lat", FloatToString(location.Latitude))
	q.Add("lng", FloatToString(location.Longitude))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&ss)
	location.Sunrise = ss.Results.Sunrise
	location.Sunset = ss.Results.Sunset
	location.DayLength = ss.Results.DayLength
}

func main() {
	ipAddress := GetIP()
	location := GetCoordinatesFromIP(ipAddress)
	location.GetSunriseSunset()
	fmt.Printf("sunrise: %s UTC\nsunset: %s UTC\nday length: %s\n", location.Sunrise, location.Sunset, location.DayLength)
}
