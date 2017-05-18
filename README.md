# daylight

This little command line util tells you the sunrise and sunset time of your
current location by grabbing the latitude and longitude based on your IP
address (hence, it may not work if you're behind a VPN, but there are plans to
be able to input your IANA timezone in the future). As you might imagine, its
usefulness is inversely proportional to how close you are to the equator.

This was written as a fun intro project to Go. :)

Inspired by [`wttr.in`](https://github.com/chubin/wttr.in).


### Usage

Compile the binary and run the program:

    $ make
    $ ./daylight
    Oslo, Norway
    sunrise: 4:35 AM
    sunset: 9:51 PM


### APIs and resources used

* <http://checkip.amazonaws.com/>
* <https://sunrise-sunset.org/api>
* <https://freegeoip.net>


### Links

* <https://www.reddit.com/r/golang/comments/3hdzza/ip_geolocation_in_go/cu84x1g/>
