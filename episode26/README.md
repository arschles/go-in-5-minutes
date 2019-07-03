# Consuming a REST API in Go

Go in 5 Minutes, episode 26.

In this screencast, we're going to build a command line client to consume the awesome [Dark Sky API](https://darksky.net/dev/docs).

We'll be using [cobra](https://github.com/spf13/cobra) to build our command line client, and I did a previous episode on that package. If you haven't seen [episode 18](https://www.goin5minutes.com/screencast/episode_18_cli_with_cobra/), you might want to go review that before you look at this one.

Instead of using an already-built Dark Sky API client (there are a [few](https://darksky.net/dev/docs/libraries) for Go), we're going to build our own client according to the API documentation. For that, we'll use the awesome [gorequest](https://github.com/parnurzeal/gorequest) package.

# Outline

1. Quick primer on Cobra
1. Quick primer on gorequest
1. Let's check out the code!

# Show Notes

- [Cobra CLI Package](https://github.com/spf13/cobra)
- [Cobra Code Generation CLI](https://github.com/spf13/cobra/blob/master/cobra/README.md)
- [GitHub client library](https://godoc.org/github.com/google/go-github/github)
