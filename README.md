# CDN Geolocation  [<img alt="resolve.rs logo" src="assets/favicon.svg" height="90" align="right" />](https://resolve.rs/)

Server that determines your physical location by looking at headers sent from various hosting providers and content delivery networks (CDNs).

<a href="https://ae-geo.redirect2.me/">Try it with Google AppEngine</a>.

<a href="https://cf-geo.redirect2.me/">Try it with Cloudflare</a>.

<a href="https://resolve.rs/ip/geolocation.html">Comparison of geolocation providers</a>.

## How it works

This application is just a simple app that shows various HTTP header values.

Applications running on [Google AppEngine](https://cloud.google.com/appengine) get some additional HTTP headers that pinpoint the client's location: [Official documentation](https://cloud.google.com/appengine/docs/standard/go/reference/request-response-headers)

Applications running behind [Cloudflare's CDN](https://www.cloudflare.com/) gets some additional HTTP headers, including one that indicates the country,  [Official documentation](https://support.cloudflare.com/hc/en-us/articles/200168236-Configuring-Cloudflare-IP-Geolocation)

## Contributions

Contributions are welcome!  If you know of any other similar CDNs/services, let me know & I will add them in!

## API

There is a simple JSON/JSONP API that is free for light, non-commercial use.  This is such a trivial application that you should run your own copy (or make your own  version) for anything serious.  Both AppEngine and Cloudflare have generous free plans (which is what I'm using).

Send a `callback` parameter to get JSONP instead of JSON.

`https://cf-geo.redirect2.me/api/cloudflare.json` for Cloudflare results

`https://ae-geo.redirect2.me/api/appengine.json` for AppEngine results

## License

[GNU Affero General Public License v3.0](LICENSE.txt)

## Credits

[![Cloudflare](https://www.vectorlogo.zone/logos/cloudflare/cloudflare-ar21.svg)](https://www.cloudflare.com/ "CDN")
[![Git](https://www.vectorlogo.zone/logos/git-scm/git-scm-ar21.svg)](https://git-scm.com/ "Version control")
[![Github](https://www.vectorlogo.zone/logos/github/github-ar21.svg)](https://github.com/ "Code hosting")
[![golang](https://www.vectorlogo.zone/logos/golang/golang-ar21.svg)](https://golang.org/ "Programming language")
[![Google AppEngine](https://www.vectorlogo.zone/logos/google_appengine/google_appengine-ar21.svg)](https://cloud.google.com/appengine/ "Hosting")
[![svgrepo](https://www.vectorlogo.zone/logos/svgrepo/svgrepo-ar21.svg)](https://www.svgrepo.com/svg/185727/map-position "favicon")

* togo `go run github.com/flazz/togo --pkg=main --name=faviconIco --input=assets/favicon.ico`
