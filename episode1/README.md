# Building a RESTful API with net/http

[![Watch The Screencast](https://www.goin5minutes.com/img/watch-screencast.svg)](https://www.goin5minutes.com/screencast/episode_1_building_restful_api_using_only_std_lib/)

Go in 5 Minutes, episode 1.

This screencast shows how to build a non-trivial HTTP API using nothing but [`net/http`](https://godoc.org/net/http).

Screencast video:
https://www.goin5minutes.com/screencast/episode_1_building_restful_api_using_only_std_lib/

## Administrivia

1. [Request a Screencast](https://github.com/arschles/go-in-5-minutes#request-a-screencast)
2. Around the web
 - [The Changelog](http://email.changelog.com/t/ViewEmail/t/D4E0966AA0002771)
 - [Reddit](https://www.reddit.com/r/golang/comments/3mpbyh/weekly_5_minute_screencast_for_gophers/)
 - [Golang Bridge](https://forum.golangbridge.org/t/new-screencast-for-gophers/124)

## Outline

1. Intro: orthogonal concerns
  - Routes
  - Handlers
  - Server
2. Using [ServeMux](https://godoc.org/net/http#ServeMux) to register paths
3. Writing handlers
