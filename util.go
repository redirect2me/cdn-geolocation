package main

import (
	"net"
	"net/http"
	"strings"
)

func getIpAddress(r *http.Request) string {
	ip := strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]

	if ip == "" {
		directIp, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			//LATER: log? try to parse?
			ip = directIp
		} else {
			ip = directIp
		}
	}

	return ip
}

func getFlatHeaders(r *http.Request, prefix string) map[string]string {

	result := make(map[string]string)
	for name, values := range r.Header {
		if strings.HasPrefix(name, prefix) {
			result[name] = values[0]
		}
	}

	return result
}

func getHeaders(r *http.Request) map[string][]string {

	result := make(map[string][]string)
	for name, values := range r.Header {
		result[name] = values
	}

	return result
}
