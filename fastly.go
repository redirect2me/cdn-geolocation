package main

import (
	"fmt"
	"html"
	"net/http"
	"time"
)

func fastlyRootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf8")
		w.Write([]byte(`<html>
	<head>
    <meta charset="utf-8">
        <title>Fastly Geolocation - Resolve.rs</title>
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/kognise/water.css@latest/dist/light.min.css" />
	</head>
    <body>
        <h1>
            <img alt="Resolve.rs geolocation logo" src="favicon.svg" style="height:2.2em;vertical-align:middle;" />
            Fastly Geolocation
        </h1>
        <p>
			Determine your real (physical) location based on your IP address, powered by Fastly.
        </p>
		<p>
            Your IP address:`))

		fmt.Fprintf(w, "%s", getIpAddress(r))
		fmt.Fprintf(w, "</p><p>")
		fmt.Fprintf(w, "Country: %s<br/>", html.EscapeString(getHeader(r, "X-Fastly-Geo-Country-Name-Utf8", "(none)")))
		fmt.Fprintf(w, "Region: %s<br/>", html.EscapeString(getHeader(r, "X-Fastly-Geo-Region-Utf8", "(none)")))
		fmt.Fprintf(w, "City: %s<br/>", html.EscapeString(getHeader(r, "X-Fastly-Geo-City-Utf8", "(none)")))
		fmt.Fprintf(w, "Latitude/Longitude: %s,%s<br/>",
			html.EscapeString(getHeader(r, "X-Fastly-Geo-Latitude", "(none)")),
			html.EscapeString(getHeader(r, "X-Fastly-Geo-Longitude", "(none)")))
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

type fastlyApiResponse struct {
	Success   bool              `json:"success"`
	Message   string            `json:"message"`
	Timestamp string            `json:"timestamp"`
	IpAddress string            `json:"ip"`
	Country   string            `json:"country"`
	Text      string            `json:"text"`
	Region    string            `json:"region"`
	City      string            `json:"city"`
	Latitude  string            `json:"latitude"`
	Longitude string            `json:"longitude"`
	Raw       map[string]string `json:"raw"`
}

func fastlyApiHandler(w http.ResponseWriter, r *http.Request) {
	result := fastlyApiResponse{}
	result.Timestamp = time.Now().UTC().Format(time.RFC3339)
	result.IpAddress = getIpAddress(r)
	result.Raw = getFlatHeaders(r, "X-Fastly-Geo-")
	for k, v := range getFlatHeaders(r, "Fastly-") {
		result.Raw[k] = v
	}
	result.Raw["X-Forwarded-For"] = r.Header.Get("X-Forwarded-For")

	result.Success = true
	result.Message = "Free for light, non-commercial use"
	result.Country = getHeader(r, "X-Fastly-Geo-Country-Code", "(not set)")
	city := getHeader(r, "X-Fastly-Geo-City-Utf8", "(not set)")
	region := getHeader(r, "X-Fastly-Geo-Region-Utf8", "(not set)")
	country := getHeader(r, "X-Fastly-Geo-Country-Name-Utf8", "(not set)")
	result.Text = fmt.Sprintf("%s, %s, %s", city, region, country)
	result.Latitude = r.Header.Get("X-Fastly-Geo-Latitude")
	result.Longitude = r.Header.Get("X-Fastly-Geo-Longitude")
	write_with_callback(w, r, result)

}
