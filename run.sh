#!/bin/bash

export LASTMOD=$(date -u)
export COMMIT=local

go run server.go appengine.go aws.go cloudflare.go jsonp.go faviconIco.go status.go util.go --port=4000 --verbose --awshost=localhost:4000
