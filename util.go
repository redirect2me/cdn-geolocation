package main

import (
	"net/http"
	"strings"
)

func getIpAddress(r *http.Request) string {
	return strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]
}

