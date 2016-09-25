+++
date = "2015-11-09T13:32:59-08:00"
title = "Writing Unit Tests Against HTTP Handlers"
type = "screencast"

+++

_Episode 6_

We show how to write unit tests against HTTP handlers. Using an example API server that proxies the [Github API](https://developer.github.com/v3/) using [Google's go-github client library](https://godoc.org/github.com/google/go-github/github), we outline and show examples of two methods for testing.
<!--more-->

These two methods are:

1. Direct testing using [`net/http/httptest.ResponseRecorder`](https://godoc.org/net/http/httptest#ResponseRecorder)
2. Integrated testing using [`github.com/arschles/testsrv`](https://godoc.org/github.com/arschles/testsrv)

Note that the second method is not strictly unit testing because it tests the router and handler at the same time. This technique is necessary in some cases because some routers encourage or require tight coupling with their handlers. The [example code](https://github.com/arschles/go-in-5-minutes/tree/master/episode6) shows such an example with [Gorilla Mux](http://godoc.org/github.com/gorilla/mux), a favorite library of this screencast.

Note also that I wrote the [testsrv](https://godoc.org/github.com/arschles/testsrv) library. It is a convenience wrapper on top of [`net/http/httptest.Server`](https://godoc.org/net/http/httptest#Server) and example code shows how to use it. [Pull requests](https://github.com/arschles/testsrv/pulls) and [issues](https://github.com/arschles/testsrv/issues) are welcome on `testsrv`.

<iframe
  class="ytplayer"
  type="text/html"
  width="640"
  height="390"
  src="https://www.youtube.com/embed/YmbbmyxSlcg?autoplay=0&origin=https://www.goin5minutes.com"
  frameborder="0"
></iframe>

Check out the example code [on Github](https://github.com/arschles/go-in-5-minutes/tree/master/episode6).
