package main

import (
	"flag"
	"github.com/ysim/daylight"
)

func main() {
	// If the `date` flag is omitted, it will default to "today"
	date := flag.String("date", "today", "Get information for this date")
	flag.Parse()

	ipAddress := daylight.GetIP()
	location := daylight.GetLocationInfoFromIP(ipAddress)
	location.GetSunriseSunset(*date)
	location.GetLocalizedSunriseSunset()
	location.Display()
}
