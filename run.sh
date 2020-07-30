#!/bin/bash

export LASTMOD=$(date -u)
export COMMIT=local

go run server.go appengine.go cloudflare.go jsonp.go faviconIco.go status.go --port=4000 --verbose
