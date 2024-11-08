package main

import (
	//	"bytes"
	//	"context"
	//	"encoding/json"
	"flag"

	//	"io/ioutil"
	"log"
	//	"net"
	"net/http"
	//	"net/url"
	"os"
	"strconv"
	//	"strings"
)

var (
	verbose        = flag.Bool("verbose", true, "verbose logging")
	aeHostname     = flag.String("aehost", "ae-geo.redirect2.me", "hostname for AppEngine")
	cfHostname     = flag.String("cfhost", "cf-geo.redirect2.me", "hostname for Cloudflare")
	awsHostname    = flag.String("awshost", "origin-aws-geo.redirect2.me", "origin hostname for AWS CloudFront (not the actual hostname)")
	bunnyHostname  = flag.String("bunnyhost", "origin-bunny-geo.redirect2.me", "origin hostname for Bunny.net (not the actual hostname)")
	fastlyHostname = flag.String("fastlyhost", "cdn-geo.global.ssl.fastly.net", "hostname for Fastly")

	logger = log.New(os.Stdout, "R2ME-GEO: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
)

func robotsTxtHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf8")
	w.Write([]byte(`#
# robots.txt for redirect2.me's geolocation servers
#
#
# not much here, but feel free to index it
#

User-Agent: *
Allow: /`))
}

func faviconSvgHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf8")
	w.Write([]byte(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 400 400"><path d="M320 109l76 28v244l-76-28V109z" fill="#d3f5ff"/><path d="M316 109l-76 28v244l76-28V109z" fill="#b7efff"/><path d="M160 109l76 28v244l-76-28V109z" fill="#d3f5ff"/><path d="M156 109l-76 28v244l76-28V109z" fill="#b7efff"/><path d="M4 109l72 28v244L4 353V109z" fill="#d3f5ff"/><g fill="#e5f9ff"><path d="M59.2 374.6L76 345.8V251L14.8 357l44.4 17.6zM236 169.8L160 301v52l76 28V169.8z"/></g><path d="M156 355.4V223L78.4 357v23.6l77.6-25.2z" fill="#a1e8fd"/><g fill="#e5f9ff"><path d="M76 220.2L4 344.6v-28.4l72-124.8v28.8zM396 268.6l-53.6 92.8-13.6-4.8 67.2-116v28z"/></g><path d="M317.2 251.4L240 384.6v-54.4L317.2 197v54.4z" fill="#a1e8fd"/><path d="M240 212.6c20-34 42-65.2 47.2-92.8L240 137v75.6z" fill="#ebd0ca"/><path d="M320 288.2l76-131.6V137l-76-28v179.2z" fill="#e5f9ff"/><g fill="#64bce1"><path d="M8 117v228c0 3.2 2.8 7.6 6 9.2l61.6 24.4c2.4.8 6.8.8 8.8 0l65.2-26c6-2.4 14.8-2.4 20.8 0l65.2 26c2.4.8 6.8.8 8.8 0l65.2-26c6-2.4 14.8-2.4 20.8 0l57.6 23.2c2.8 1.2 4 .4 4-2.8V145c0-3.2-2.8-7.6-6-9.2l-61.6-24.4c-2.4-.8-6.8-.8-8.8 0l-65.2 26c-6 2.4-14.8 2.4-20.8 0l-65.2-26c-2.4-.8-6.8-.8-8.8 0l-65.2 26c-6 2.4-14.8 2.4-20.8 0L12 114.2c-3.2-1.2-4-.4-4 2.8zm-8 0c0-8.8 6.8-13.2 14.8-10l57.6 23.2c4 1.6 10.8 1.6 14.8 0l65.2-26c4-1.6 10.8-1.6 14.8 0l65.2 26c4 1.6 10.8 1.6 14.8 0l65.2-26c4-1.6 10.8-1.6 14.8 0l61.6 24.4c6 2.4 11.2 9.6 11.2 16.4v228c0 8.8-6.8 13.2-14.8 10l-57.6-23.2c-4-1.6-10.8-1.6-14.8 0l-65.2 26c-4 1.6-10.8 1.6-14.8 0l-65.2-26c-4-1.6-10.8-1.6-14.8 0l-65.2 26c-4 1.6-10.8 1.6-14.8 0l-61.6-24.4C4.8 359 0 351.8 0 345V117z"/><path d="M76 137h4v244h-4z"/></g><path d="M152 136.2L80 260.6v-28.4l72-124.8v28.8z" fill="#a1e8fd"/><path fill="#64bce1" d="M156 109h4v244h-4z"/><path d="M317.2 167.4L240 300.6v-30.4L317.2 137v30.4z" fill="#a1e8fd"/><g fill="#64bce1"><path d="M236 137h4v244h-4zM316 109h4v244h-4z"/></g><path d="M236 219.4c-9.6 16.4-18 33.6-24 51.6-1.6 4.4-1.6 4.8-4 6-2.4-.8-2.4-1.6-4-6-10-30-28-58.4-44-84.8V109l29.2 10.8c4.4 5.6 11.2 9.2 18.8 9.2 1.6 0 3.2 0 4.8-.4L236 137v82.4z" fill="#e5cec8"/><path d="M200 13c-44 0-80 36-80 80 0 44.8 55.6 99.6 76 162 1.6 4.8 1.6 5.2 4 6 2.4-.8 2.4-1.6 4-6 20.4-62 76-117.2 76-162 0-44-36-80-80-80zm-.8 98.8c-12.4 0-22.8-10-22.8-22.8s10-22.8 22.8-22.8c12.4 0 22.8 10 22.8 22.8s-10 22.8-22.8 22.8z" fill="#f88765"/><path d="M200 253.8c0 .4 0 .4 0 0 7.2-20.8 17.2-40.4 34.8-69.6 0-.4 12-20 15.2-25.2 5.6-9.2 9.6-16.8 13.2-23.6C271.6 118.6 276 105 276 93c0-42-34-76-76-76s-76 34-76 76c0 12 4.4 25.6 12.8 42 3.6 6.8 7.6 14.4 13.2 23.6 3.2 5.6 15.2 24.8 15.2 25.2 17.6 29.6 27.6 48.8 34.8 70zm-4 1.2c-20.4-62.4-76-117.2-76-162 0-44 36-80 80-80s80 36 80 80c0 44.8-55.6 100-76 162-1.6 4.4-1.6 4.8-4 6-2.4-.8-2.4-1.6-4-6zm4-146c11.2 0 20-8.8 20-20s-8.8-20-20-20-20 8.8-20 20 8.8 20 20 20zm0 4c-13.2 0-24-10.8-24-24s10.8-24 24-24 24 10.8 24 24-10.8 24-24 24z" fill="#e85a2f"/></svg>`))
}

func faviconIcoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/vnd.microsoft.icon")
	w.Write(faviconIco)
}

func getHeader(r *http.Request, key string, defaultValue string) string {
	retVal := r.Header.Get(key)
	if retVal == "" {
		retVal = defaultValue
	}

	return retVal
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	host := getHost(r)
	if r.URL.Path[1:] == "" {
		if host == *aeHostname {
			appengineRootHandler(w, r)
		} else if host == *cfHostname {
			cloudflareRootHandler(w, r)
		} else if host == *fastlyHostname {
			fastlyRootHandler(w, r)
		} else if host == *awsHostname {
			awsRootHandler(w, r)
		} else if host == *bunnyHostname {
			bunnyRootHandler(w, r)
		} else {
			logger.Printf("WARN: unknown host '%s'\n", host)
			http.Redirect(w, r, "https://github.com/redirect2me/cdn-geolocation", http.StatusTemporaryRedirect)
		}
	} else {
		logger.Printf("WARN: unknown page '%s' from '%s' on '%s'\n", r.URL.Path, getHeader(r, "Referer", "(no referrer)"), host)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

func main() {

	var defaultPort, portErr = strconv.Atoi(os.Getenv("PORT"))
	if portErr != nil {
		defaultPort = 4000
	}
	var port = flag.Int("port", defaultPort, "port to listen on")

	flag.Parse()

	if *verbose {
		logger.Printf("DEBUG: running in verbose mode\n")
	}

	http.HandleFunc("/status.json", statusHandler)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/robots.txt", robotsTxtHandler)
	http.HandleFunc("/favicon.ico", faviconIcoHandler)
	http.HandleFunc("/favicon.svg", faviconSvgHandler)
	http.HandleFunc("/headers.html", headersHandler)

	http.HandleFunc("/api/appengine.json", appengineApiHandler)
	http.HandleFunc("/api/bunny.json", bunnyApiHandler)
	http.HandleFunc("/api/cloudflare.json", cloudflareApiHandler)
	http.HandleFunc("/api/fastly.json", fastlyApiHandler)
	http.HandleFunc("/api/aws.json", awsApiHandler)
	http.HandleFunc("/api/headers.json", headersApiHandler)

	if *verbose {
		logger.Printf("INFO: running on port %d\n", *port)
	}
	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err != nil {
		logger.Panicf("ERROR: unable to listen on port %d: %s\n", *port, err)
	}
}
