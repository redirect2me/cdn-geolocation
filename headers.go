package main

import (
	//	"bytes"
	//	"context"
	//	"encoding/json"

	"fmt"
	"html"
	"net"
	"time"

	//	"io/ioutil"

	//	"net"
	"net/http"
	//	"net/url"
	//	"strings"
)

func headersHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.Write([]byte(`<html>
	<head>
    <meta charset="utf-8">
        <title>HTTP Headers - Resolve.rs</title>
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/kognise/water.css@latest/dist/light.min.css" />
	</head>
    <body>
		<h1>
			Debug page for HTTP headers
        </h1>
        <p>
            These are the HTTP headers seen by this server.
        </p>
		<p>
            Raw IP address:`))

	rawIp, _, _ := net.SplitHostPort(r.RemoteAddr)

	fmt.Fprintf(w, "%s", rawIp)
	w.Write([]byte(`<br/>Calculated IP address: `))
	fmt.Fprintf(w, "%s", getIpAddress(r))
	w.Write([]byte(`<br/>Raw Host: `))
	fmt.Fprintf(w, "%s", html.EscapeString(r.Host))
	w.Write([]byte(`<br/>Calculated Host: `))
	fmt.Fprintf(w, "%s", html.EscapeString(getHost(r)))
	w.Write([]byte(`</p>
		<table>
			<thead>
				<tr>
					<th>Name</th>
					<th>Value(s)</th>
				</tr>
			</thead>
			<tbody>
			`))

	for name, values := range r.Header {
		fmt.Fprintf(w, "<tr><td>%s</td><td>", html.EscapeString(name))

		if len(values) == 1 {
			fmt.Fprintf(w, "%s", html.EscapeString(values[0]))
		} else {
			for index, value := range values {
				fmt.Fprintf(w, "%d: %s<br/>", index, html.EscapeString(value))
			}
		}
		w.Write([]byte(`</td></tr>`))
	}

	w.Write([]byte(`   </tbody>
		</table>
        <p>
            <a href="https://github.com/redirect2me/cdn-geolocation">How this works</a>, including API details and source code!
        </p>
        <p>
            <a href="https://resolve.rs/">Resolve.rs</a>
            has more
            <a href="https://resolve.rs/tools.html#http">HTTP troubleshooting tools</a>.
        </p>
	</body>
</html>`))
}

type headersApiResponse struct {
	Success   bool                `json:"success"`
	Message   string              `json:"message"`
	Headers   map[string][]string `json:"headers"`
	Timestamp string              `json:"timestamp"`
	IpAddress string              `json:"ip"`
}

func headersApiHandler(w http.ResponseWriter, r *http.Request) {
	result := headersApiResponse{}
	result.Timestamp = time.Now().UTC().Format(time.RFC3339)
	result.IpAddress = getIpAddress(r)

	result.Success = true
	result.Message = "Free for light, non-commercial use"

	result.Headers = getHeaders(r)

	write_with_callback(w, r, result)
}
