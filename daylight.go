package daylight

import (
	"encoding/json"
	"fmt"
	"github.com/bradfitz/latlong"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	externalIPurl             = "http://checkip.amazonaws.com/"
	geoIPurl                  = "https://freegeoip.net/json/"
	sunriseSunsetUrl          = "https://api.sunrise-sunset.org/json"
	googleMapsGeocodingAPIurl = "https://maps.googleapis.com/maps/api/geocode/json"
)

var googleApiKey = os.Getenv("GOOGLE_MAPS_API_KEY")

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
	City         string
	Country      string
	Latitude     float64
	Longitude    float64
	Timezone     string
	SunriseUTC   string
	SunsetUTC    string
	DayLength    string
	SunriseLocal time.Time
	SunsetLocal  time.Time
}

type City struct {
	Name string
}

func UnpackGeocodingResultToLocation(geocode maps.GeocodingResult) Location {
	// Get the City and Country
	city := geocode.AddressComponents[0].LongName
	country := geocode.AddressComponents[len(geocode.AddressComponents)-1].LongName
	latitude := geocode.Geometry.Location.Lat
	longitude := geocode.Geometry.Location.Lng
	timezone := latlong.LookupZoneName(latitude, longitude)
	location := Location{
		City:      city,
		Country:   country,
		Latitude:  latitude,
		Longitude: longitude,
		Timezone:  timezone,
	}
	return location
}

func BuildLocation(address string, ip string) Location {
	// TODO: Error handling in the event that both or none are provided

	// If no address provided in cli, get timezone from ip address
	if address == "" {
		location := GetLocationInfoFromIP(ip)
		return location
	}

	// Alternatively, use address from user input if one was provided
	// First, look up the coordinates using the Google Maps Geocoding API
	c, err := maps.NewClient(maps.WithAPIKey(googleApiKey))
	if err != nil {
		log.Fatal(err)
	}
	request := &maps.GeocodingRequest{
		Address: address,
	}
	response, err := c.Geocode(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	location := UnpackGeocodingResultToLocation(response[0])
	// TODO: Uncomment when logging levels are figured out
	//log.Printf("%+v\n", location)
	return location
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

func ExtractCityFromTimezone(timezone string) City {
	splitTimezone := strings.Split(timezone, "/")
	city := City{splitTimezone[len(splitTimezone)-1]}
	city.Name = strings.Replace(city.Name, "_", " ", -1)
	return city
}

func GetLocationInfoFromIP(ip string) Location {
	response, err := http.Get(geoIPurl + ip)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(response.Body).Decode(&geo)
	location := Location{
		City:      geo.City,
		Country:   geo.CountryName,
		Timezone:  geo.Timezone,
		Latitude:  geo.Latitude,
		Longitude: geo.Longitude,
	}

	// Sometimes freegeoip doesn't return the city name for some reason
	if geo.City == "" {
		location.City = ExtractCityFromTimezone(geo.Timezone).Name
	}
	return location
}

func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', 7, 64)
}

func (location *Location) GetSunriseSunset(date string) {
	// Build the request
	client := &http.Client{}
	req, err := http.NewRequest("GET", sunriseSunsetUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("lat", FloatToString(location.Latitude))
	q.Add("lng", FloatToString(location.Longitude))
	q.Add("formatted", "0")
	q.Add("date", date)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&ss)
	location.SunriseUTC = ss.Results.Sunrise
	location.SunsetUTC = ss.Results.Sunset
	location.DayLength = ss.Results.DayLength
}

func StringToTime(timeInput string) (time.Time, error) {
	layout := "2006-01-02T15:04:05-07:00"
	timeParsed, err := time.Parse(layout, timeInput)
	if err != nil {
		log.Fatal("Unable to parse time:", timeInput)
	}
	//log.Printf("Parsed '%s' to '%s'\n", timeInput, timeParsed)
	return timeParsed, nil
}

func (location *Location) GetLocalizedSunriseSunset() {
	// Populates the Sunrise and Sunset fields
	// TODO: Check that SunriseUTC, SunsetUTC, and Timezone are populated
	sunriseUTCTime, _ := StringToTime(location.SunriseUTC)
	sunsetUTCTime, _ := StringToTime(location.SunsetUTC)

	// TODO: time.LoadLocation will just assume UTC if location.Timezone is nil,
	// resulting in location.Sun{rise,set}Local being in UTC time without
	// erroring
	localizedTimezone, err := time.LoadLocation(location.Timezone)
	if err != nil {
		log.Fatal("Unable to load location: %s", err)
	}
	location.SunriseLocal = sunriseUTCTime.In(localizedTimezone)
	location.SunsetLocal = sunsetUTCTime.In(localizedTimezone)
}

func FormatTimeForUser(timeInput time.Time) string {
	layout := "3:04 PM"
	formattedTime := timeInput.Format(layout)
	return formattedTime
}

func (location *Location) GetDisplayString() string {
	// Info that is printed to the screen for the user
	formattedSunriseTime := FormatTimeForUser(location.SunriseLocal)
	formattedSunsetTime := FormatTimeForUser(location.SunsetLocal)
	displayString := fmt.Sprintf("%s, %s\nsunrise: %s\nsunset: %s", location.City, location.Country, formattedSunriseTime, formattedSunsetTime)
	return displayString
}
