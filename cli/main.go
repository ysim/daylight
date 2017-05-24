package main

import (
	"flag"
	"github.com/ysim/daylight"
)

func main() {
	// If the `date` flag is omitted, it will default to "today"
	date := flag.String("date", "today", "Get information for this date")
	address := flag.String("address", "", "Get information for this address")
	flag.Parse()

	// Get a Location struct either based on IP or address input
	// It should include the City, Country, Latitude/Longitude, and Timezone
	location := daylight.BuildLocation(*address, "")

	// Now populate the location struct with SunriseUTC and SunsetUTC
	location.GetSunriseSunset(*date)

	// ...and SunriseLocal/SunsetLocal...
	location.GetLocalizedSunriseSunset()

	// Now display the info.
	location.Display()
}
