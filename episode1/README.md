# Building a RESTful API with Gorilla Mux

Go in 5 Minutes, episode 1.

This screencast shows off some of the features of Gorilla Mux by implementing a key/value store
behind a RESTful API.

## Administrivia

1. [Request a Screencast](https://github.com/arschles/go-in-5-minutes#request-a-screencast)
2. Around the web
 - [The Changelog](http://email.changelog.com/t/ViewEmail/t/D4E0966AA0002771)
 - [Reddit](https://www.reddit.com/r/golang/comments/3mpbyh/weekly_5_minute_screencast_for_gophers/)
 - [Golang Bridge](https://forum.golangbridge.org/t/new-screencast-for-gophers/124)

## Outline

1. Intro: orthogonal concerns
  - ServeMux
  - Handlers
  - Server
2. Using [Routers](https://godoc.org/github.com/gorilla/mux#Router) to register paths
3. Getting [path vars](https://godoc.org/github.com/gorilla/mux#Vars)
4. Serving
