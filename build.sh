#!/bin/bash


go build \
	-ldflags "-X main.COMMIT=$(git rev-parse --short HEAD) -X main.LASTMOD=$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
	server.go appengine.go aws.go cloudflare.go fastly.go faviconIco.go headers.go jsonp.go status.go util.go
