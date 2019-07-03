# Consuming a REST API in Go

Go in 5 Minutes, episode 26.

In this screencast, we're going to build a command line client to consume the awesome [Dark Sky API](https://darksky.net/dev/docs).

We'll be using [cobra](https://github.com/spf13/cobra) to build our command line client, and I did a previous episode on that package. If you haven't seen [episode 18](https://www.goin5minutes.com/screencast/episode_18_cli_with_cobra/), you might want to go review that before you look at this one.

Instead of using an already-built Dark Sky API client (there are a [few](https://darksky.net/dev/docs/libraries) for Go), we're going to build our own client according to the API documentation to show some tips and tricks for building clients for any REST API.

In this screencast, we'll use the awesome [gorequest](https://github.com/parnurzeal/gorequest) package to help us build a DarkSky client from scratch.

# Outline

1. Quick primer on Cobra
1. Quick primer on gorequest
1. Let's check out the code!

# How to Run This Code

You'll need Go version 1.11 or above to run this code. If you have an appropriate version, simply run `go build -o darksky .` to build. 

Before you run the binary, you'll need an environment variable called `DARKSKY_API_KEY` set to your DarkSky API key (if you don't have one, get it from your [account](https://darksky.net/dev/account), or [create](https://darksky.net/dev/register) an account if you haven't already).

Then, call the binary like so, ensuring that `DARKSKY_API_KEY` is set in your environment:

```console
$ ./darksky temp --lat 45.512230 --long -122.658722
```

The `lat` and `long` flags are set to the latitude and longitude (respectively) of the location for which to get the temperature.

>The latitude and longitude in the above example are set to Portland, OR, USA. If you'd like to try another location, you can use https://www.latlong.net/

# Show Notes

- [Cobra CLI Package](https://github.com/spf13/cobra)
- [Cobra Code Generation CLI](https://github.com/spf13/cobra/blob/master/cobra/README.md)
- [GitHub client library](https://godoc.org/github.com/google/go-github/github)
