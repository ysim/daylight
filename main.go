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

type Coordinates struct {
	Latitude  float64
	Longitude float64
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

func GetCoordinatesFromIP(ip string) *Coordinates {
	response, err := http.Get(geoIPurl + ip)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(response.Body).Decode(&geo)
	log.Printf("Getting latitude and longitude for %s, %s", geo.City, geo.CountryName)
	c := &Coordinates{Latitude: geo.Latitude, Longitude: geo.Longitude}
	return c
}

func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', 7, 64)
}

func GetSunriseSunset(c *Coordinates) (string, string, string) {
	// Build the request
	client := &http.Client{}
	req, err := http.NewRequest("GET", sunriseSunsetUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("lat", FloatToString(c.Latitude))
	q.Add("lng", FloatToString(c.Longitude))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&ss)
	return ss.Results.Sunrise, ss.Results.Sunset, ss.Results.DayLength
}

func main() {
	ipAddress := GetIP()
	coordinates := GetCoordinatesFromIP(ipAddress)
	sunrise, sunset, dayLength := GetSunriseSunset(coordinates)
	fmt.Printf("sunrise: %s UTC\nsunset: %s UTC\nday length: %s\n", sunrise, sunset, dayLength)
}
