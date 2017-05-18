package main

import (
	"github.com/ysim/daylight"
)

func main() {
	ipAddress := daylight.GetIP()
	location := daylight.GetLocationInfoFromIP(ipAddress)
	location.GetSunriseSunset()
	location.GetLocalizedSunriseSunset()
	location.Display()
}
