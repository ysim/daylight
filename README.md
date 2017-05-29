# daylight

This little command line util tells you the sunrise and sunset time of your
current location by grabbing the latitude and longitude based on your IP
address (hence, it may not work if you're behind a VPN, but there are plans to
be able to input your IANA timezone in the future). As you might imagine, its
usefulness is inversely proportional to how close you are to the equator.

This was written as a fun intro project to Go. :)

Inspired by [`wttr.in`](https://github.com/chubin/wttr.in).


### Installation

Compile the binary:

    make

Install the binary somewhere on your path (I like `$HOME/bin` on OS X; YMMV):

    install bin/daylight-cli $HOME/bin


### Usage

Running the binary with no flags gives you the sunrise/sunset times where you
are (based on your IP address), today:

    $ daylight-cli
    Oslo, Norway
    sunrise: 4:35 AM
    sunset: 9:51 PM

Get the local sunrise/sunset times for another day:

    $ daylight-cli -date 2016-09-25
    Oslo, Norway
    sunrise: 7:10 AM
    sunset: 7:06 PM

Get the sunrise/sunset times for a particular place for today:

    $ daylight-cli -address Ushuaia
    Ushuaia, Argentina
    sunrise: 12:34 PM
    sunset: 8:25 PM

Get the sunrise/sunset times for a particular place for some other day:

    $ daylight-cli -address Toronto -date 2016-12-23
    Toronto, Canada
    sunrise: 7:49 AM
    sunset: 4:45 PM


### APIs and resources used

* <http://checkip.amazonaws.com/>
* <https://sunrise-sunset.org/api>
* <https://freegeoip.net>
* Google Maps Geocoding API: <https://github.com/googlemaps/google-maps-services-go> (server only)


### Links

* <https://www.reddit.com/r/golang/comments/3hdzza/ip_geolocation_in_go/cu84x1g/>
* <https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091>
