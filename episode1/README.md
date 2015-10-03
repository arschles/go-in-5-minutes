# An HTTP Server With Substance

Go in 5 Minutes, episode 1.

This screencast shows how to write a full featured RESTful HTTP server. We'll build
a simple key/value store behind an HTTP API.

## Administrivia

1. [Request a Screencast](https://github.com/arschles/go-in-5-minutes#request-a-screencast)
2. Around the web
  1. [The Changelog](https://github.com/thechangelog/ping/issues/288)
  2. [Reddit](https://www.reddit.com/r/golang/comments/3mpbyh/weekly_5_minute_screencast_for_gophers/)
  3. [Golang Bridge](https://forum.golangbridge.org/t/new-screencast-for-gophers/124)

## Outline

1. [`http.ServeMux`](https://godoc.org/net/http#ServeMux) & [Gorilla Mux Routers](https://godoc.org/github.com/gorilla/mux#Router)
2. [`http.Handler`](https://godoc.org/net/http#Handler) / [`http.HandlerFunc`](https://godoc.org/net/http#HandlerFunc)
3. [`http.Handle`](https://godoc.org/net/http#Handle) (or [`http.HandleFunc`](https://godoc.org/net/http#HandleFunc))
4. Handler dependencies
