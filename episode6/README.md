# Unit Testing HTTP Handlers

Go in 5 Minutes, episode 6.

This screencast shows two ways to quickly and effectively test your [`net/http` handlers](https://godoc.org/net/http#Handler).

## [goin5minutes.com](http://www.goin5minutes.com)

Our website is up and running! Right now, it houses a permanent list of all screencasts, but I'm planning to expand it with more auxiliary content.

__Please [sign up for our newsletter](http://www.goin5minutes.com/subscribe) to get emails when new screencasts come out.__

## Outline

1. Flash back to [episode 0](https://github.com/arschles/go-in-5-minutes/tree/master/episode0)
2. [net/http/httptest](https://godoc.org/net/http/httptest) & [testsrv](https://github.com/arschles/testsrv)
  - pure handler vs. router & handler tests
3. Example - an API server to wrap the Github API
