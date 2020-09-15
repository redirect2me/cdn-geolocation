# CDN Geolocation  [<img alt="resolve.rs logo" src="assets/favicon.svg" height="90" align="right" />](https://resolve.rs/)

Server that determines your physical location by looking at headers sent from various hosting providers and content delivery networks (CDNs).

[Try it with AWS CloudFront](https://aws-geo.redirect2.me/)

[Try it with Cloudflare](https://cf-geo.redirect2.me/)

[Try it with Fastly](http://fastly-geo.redirect2.me/)

[Try it with Google AppEngine](https://ae-geo.redirect2.me/)

[Comparison of geolocation providers](https://resolve.rs/ip/geolocation.html)

## How it works

This application is just a simple app that shows various HTTP header values.

Applications running behind [AWS CloudFront](https://aws.amazon.com/cloudfront/) gets some additional HTTP headers, including one that indicates the country,  [Official documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/RequestAndResponseBehaviorCustomOrigin.html#request-custom-headers-behavior)

Applications running behind [Cloudflare's CDN](https://www.cloudflare.com/) gets some additional HTTP headers, including one that indicates the country,  [Official documentation](https://support.cloudflare.com/hc/en-us/articles/200168236-Configuring-Cloudflare-IP-Geolocation)

Applications running behind [Fastly](https://www.fastly.com/) can set additional headers with geolocation data: [Official documentation](https://developer.fastly.com/reference/vcl/variables/geolocation/)

Applications running on [Google AppEngine](https://cloud.google.com/appengine) get some additional HTTP headers that pinpoint the client's location: [Official documentation](https://cloud.google.com/appengine/docs/standard/go/reference/request-response-headers)

## Contributions

Contributions are welcome!  If you know of any other similar CDNs/services, let me know & I will add them in!

## API

There is a simple JSON/JSONP API that is free for light, non-commercial use.  This is such a trivial application that you should run your own copy (or make your own  version) for anything serious.  Both AppEngine and Cloudflare have generous free plans (which is what I'm using).

Send a `callback` parameter to get JSONP instead of JSON.

`https://aws-geo.redirect2.me/api/aws.json` for AWS CloudFront results

`https://cf-geo.redirect2.me/api/cloudflare.json` for Cloudflare results

`https://ae-geo.redirect2.me/api/appengine.json` for Google AppEngine results

## License

[GNU Affero General Public License v3.0](LICENSE.txt)

## Credits

[![AWS](https://www.vectorlogo.zone/logos/amazon_aws/amazon_aws-ar21.svg)](https://aws.amazon.com/ "CDN and Geolocation")
[![Cloudflare](https://www.vectorlogo.zone/logos/cloudflare/cloudflare-ar21.svg)](https://www.cloudflare.com/ "CDN and Geolocation")
[![Git](https://www.vectorlogo.zone/logos/git-scm/git-scm-ar21.svg)](https://git-scm.com/ "Version control")
[![Github](https://www.vectorlogo.zone/logos/github/github-ar21.svg)](https://github.com/ "Code hosting")
[![golang](https://www.vectorlogo.zone/logos/golang/golang-ar21.svg)](https://golang.org/ "Programming language")
[![Google AppEngine](https://www.vectorlogo.zone/logos/google_appengine/google_appengine-ar21.svg)](https://cloud.google.com/appengine/ "Hosting and Geolocation")
[![svgrepo](https://www.vectorlogo.zone/logos/svgrepo/svgrepo-ar21.svg)](https://www.svgrepo.com/svg/185727/map-position "favicon")
[![water.css](https://www.vectorlogo.zone/logos/netlifyapp_watercss/netlifyapp_watercss-ar21.svg)](https://watercss.netlify.app/ "Classless CSS")

* togo `go run github.com/flazz/togo --pkg=main --name=faviconIco --input=assets/favicon.ico`
