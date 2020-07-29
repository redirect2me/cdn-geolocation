package main

import (
	"fmt"
	"html"
	"net/http"
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
	</head>
    <body>
        <h1>
            <img alt="Resolve.rs geolocation logo" src="favicon.svg" style="height:2.2em;vertical-align:middle;" />
            AppEngine Geolocation
        </h1>
        <p>
            Determine which location based on your IP address, powered by Google AppEngine.
        </p>
		<p>
            Your Location:<br/>`))

		fmt.Fprintf(w, "Country: %s<br/>", html.EscapeString(getHeader(r, "X-Appengine-Country", "(none)")))
		fmt.Fprintf(w, "Region: %s<br/>", html.EscapeString(getHeader(r, "X-Appengine-Region", "(none)")))
		fmt.Fprintf(w, "City: %s<br/>", html.EscapeString(getHeader(r, "X-Appengine-City", "(none)")))
		fmt.Fprintf(w, "Latitude/Longitude: %s<br/>", html.EscapeString(getHeader(r, "X-Appengine-CityLatLong", "(none)")))
		//LATER: hyperlink to map
		w.Write([]byte(`</p>
        <p>
            <a href="https://github.com/redirect2me/appengine-geolocation">How this works</a>, including API details and source code!
        </p>
        <p>
            <a href="https://resolve.rs/">resolve.rs</a>
            has more
            <a href="https://resolve.rs/tools.html">diagnostic tools</a>.
            including a
            <a href="https://resolve.rs/ip/geolocation.html">comparison of different geolocation APIs</a>.
        </p>
		<script>
	</body>
</html>`))
	} else {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

type ApiResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Country   string `json:"country"`
	Region    string `json:"region"`
	City      string `json:"city"`
	LatLng    string `json:"latlng"`
}

func appengineApiHandler(w http.ResponseWriter, r *http.Request) {
	result := ApiResponse{}
	result.Timestamp = time.Now().UTC().Format(time.RFC3339)

	result.Success = true
	result.Message = "Free for light, non-commercial use"
	result.Country = getHeader(r, "X-Appengine-Country", "(not set)")
	result.City = getHeader(r, "X-Appengine-City", "(not set)")
	result.Region = getHeader(r, "X-Appengine-Region", "(not set)")
	result.LatLng = getHeader(r, "X-Appengine-CityLatLong", "(not set)")
	write_with_callback(w, r, result)

}
