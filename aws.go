package main

import (
	//	"bytes"
	//	"context"
	//	"encoding/json"

	"fmt"
	"html"
	"time"

	//	"io/ioutil"

	//	"net"
	"net/http"
	//	"net/url"
	//	"strings"
)

func awsRootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf8")
		w.Write([]byte(`<html>
	<head>
    <meta charset="utf-8">
        <title>Amazon AWS CloudFront Geolocation - Resolve.rs</title>
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/kognise/water.css@latest/dist/light.min.css" />
	</head>
    <body>
        <h1>
            <img alt="Resolve.rs geolocation logo" src="favicon.svg" style="height:2.2em;vertical-align:middle;" />
            AWS CloudFront Geolocation
        </h1>
        <p>
            Determine your real (physical) location based on your IP address, powered by AWS CloudFront.
        </p>
		<p>
            Your IP address:`))

		fmt.Fprintf(w, "%s", getIpAddress(r))
		fmt.Fprintf(w, "</p><p>")

		fmt.Fprintf(w, "Country: %s<br/>", html.EscapeString(getHeader(r, "CloudFront-Viewer-Country-Name", "(none)")))
		fmt.Fprintf(w, "Region: %s<br/>", html.EscapeString(getHeader(r, "CloudFront-Viewer-Country-Region-Name", "(none)")))
		fmt.Fprintf(w, "City: %s<br/>", html.EscapeString(getHeader(r, "CloudFront-Viewer-City", "(none)")))
		fmt.Fprintf(w, "Latitude/Longitude: %s,%s<br/>",
			html.EscapeString(getHeader(r, "CloudFront-Viewer-Latitude", "(none)")),
			html.EscapeString(getHeader(r, "CloudFront-Viewer-Longitude", "(none)")))
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

type awsApiResponse struct {
	Success   bool              `json:"success"`
	Message   string            `json:"message"`
	Timestamp string            `json:"timestamp"`
	IpAddress string            `json:"ip"`
	Country   string            `json:"country"`
	Text      string            `json:"text"`
	Latitude  string            `json:"latitude"`
	Longitude string            `json:"longitude"`
	Raw       map[string]string `json:"raw"`
}

func awsApiHandler(w http.ResponseWriter, r *http.Request) {
	result := awsApiResponse{}
	result.Timestamp = time.Now().UTC().Format(time.RFC3339)
	result.IpAddress = getIpAddress(r)
	result.Raw = getFlatHeaders(r, "Cloudfront-")
	result.Raw["X-Forwarded-For"] = r.Header.Get("X-Forwarded-For")

	result.Success = true
	result.Message = "Free for light, non-commercial use"
	result.Country = getHeader(r, "CloudFront-Viewer-Country", "(not set)")
	city := getHeader(r, "CloudFront-Viewer-City", "(not set)")
	region := getHeader(r, "CloudFront-Viewer-Country-Region-Name", "(not set)")
	country := getHeader(r, "CloudFront-Viewer-Country-Name", "(not set)")
	result.Text = fmt.Sprintf("%s, %s, %s", city, region, country)
	result.Latitude = r.Header.Get("CloudFront-Viewer-Latitude")
	result.Longitude = r.Header.Get("CloudFront-Viewer-Longitude")
	write_with_callback(w, r, result)
}
