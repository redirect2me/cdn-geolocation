# AppEngine Geolocation  [<img alt="resolve.rs logo" src="assets/favicon.svg" height="90" align="right" />](https://resolve.rs/)

Server that determines your physical location, powered by Google AppEngine.  <a href="">Try it!</a>.  <a href="https://resolve.rs/ip/geolocation.html">Compare geolocation providers</a>.


## How it works

Applications running on [Google AppEngine](https://cloud.google.com/appengine) get some additional HTTP headers that pinpoint the client's location: [Official documentation](https://cloud.google.com/appengine/docs/standard/go/reference/request-response-headers)

This application is just a simple app that shows those header values.

## License

[GNU Affero General Public License v3.0](LICENSE.txt)

## Credits

golang
appengine
git
github

https://www.svgrepo.com/svg/185727/map-position
go run github.com/flazz/togo --pkg=main --name=faviconIco --input=assets/favicon.ico
