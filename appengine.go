package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func appengineRootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf8")
		w.Write([]byte(`<html>
	<head>
    <meta charset="utf-8">
        <title>AppEngine Geolocation - Resolve.rs</title>
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/kognise/water.css@latest/dist/light.min.css" />
		<link rel="icon" type="image/svg+xml" href="/favicon.svg" />
	</head>
    <body>
        <h1>
            <img alt="Resolve.rs geolocation logo" src="favicon.svg" style="height:2.2em;vertical-align:middle;" />
            AppEngine Geolocation
        </h1>
        <p>
			Determine your real (physical) location based on your IP address, powered by Google AppEngine.
        </p>
		<p>
            Your IP address:`))

		fmt.Fprintf(w, "%s", getIpAddress(r))
		fmt.Fprintf(w, "</p><p>")
		fmt.Fprintf(w, "Country: %s<br/>", html.EscapeString(getHeader(r, "X-Appengine-Country", "(none)")))
		fmt.Fprintf(w, "Region: %s<br/>", html.EscapeString(getHeader(r, "X-Appengine-Region", "(none)")))
		fmt.Fprintf(w, "City: %s<br/>", html.EscapeString(getHeader(r, "X-Appengine-City", "(none)")))
		fmt.Fprintf(w, "Latitude/Longitude: %s<br/>", html.EscapeString(getHeader(r, "X-Appengine-CityLatLong", "(none)")))
		//LATER: hyperlink to map
		w.Write([]byte(`</p>
        <p>
            <a href="https://github.com/redirect2me/cdn-geolocation">How this works</a>, including API details and source code!
        </p>
        <p>
            <a href="https://resolve.rs/">Resolve.rs</a>
            has more
            <a href="https://resolve.rs/tools.html">diagnostic tools</a>.
            including a
            <a href="https://resolve.rs/ip/geolocation.html">comparison of different geolocation APIs</a>.
        </p>
	</body>
</html>`))
	} else {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

type appengineApiResponse struct {
	Success   bool              `json:"success"`
	Message   string            `json:"message"`
	Timestamp string            `json:"timestamp"`
	IpAddress string            `json:"ip"`
	Country   string            `json:"country"`
	Text      string            `json:"text"`
	Region    string            `json:"region"`
	City      string            `json:"city"`
	Latitude  float32           `json:"latitude"`
	Longitude float32           `json:"longitude"`
	Raw       map[string]string `json:"raw"`
}

func appengineApiHandler(w http.ResponseWriter, r *http.Request) {
	result := appengineApiResponse{}
	result.Timestamp = time.Now().UTC().Format(time.RFC3339)
	result.IpAddress = getIpAddress(r)
	result.Raw = getFlatHeaders(r, "X-Appengine-")

	result.Success = true
	result.Message = "Free for light, non-commercial use"
	result.Country = getHeader(r, "X-Appengine-Country", "(not set)")
	city := getHeader(r, "X-Appengine-City", "(not set)")
	region := getHeader(r, "X-Appengine-Region", "(not set)")
	result.Text = fmt.Sprintf("%s, %s, %s", city, region, result.Country)
	latlng := r.Header.Get("X-Appengine-CityLatLong")
	if latlng != "" {
		comma := strings.Index(latlng, ",")
		if comma != -1 {
			latitude, latErr := strconv.ParseFloat(latlng[0:comma], 32)
			if latErr != nil {
				logger.Printf("ERROR: unable to convert '%s' to float: %s", latlng[0:comma], latErr)
			} else {
				result.Latitude = float32(latitude)
			}
			longitude, lngErr := strconv.ParseFloat(latlng[comma+1:len(latlng)], 32)
			if lngErr != nil {
				logger.Printf("ERROR: unable to convert '%s' to float: %s", latlng[0:comma], lngErr)
			} else {
				result.Longitude = float32(longitude)
			}
		}
	}
	write_with_callback(w, r, result)

}
