+++
type = "screencast"
title = "Using Modules with the Athens Project!"
teaser = "Module API servers are a huge boost for Go dependencies. Check out why!"
author = "Aaron Schlesinger"
date = "2019-08-09T13:10:24-07:00"
+++

_Episode 28_

In [episode 27](https://www.goin5minutes.com/screencast/episode_27_intro_to_modules/), we talked about how to use [Go Modules](https://github.com/golang/go/wiki/Modules) to manage dependencies, but there's more to the story!

Check out how module servers and Athens fit into the dependencies ecosystem.

<!--more-->

Go modules come with module servers that you can download your dependencies from too! The servers themselves are pretty cool, but here's why you should consider using one:

- Downloading dependencies will almost always be faster (sometimes up to 4x faster)
  - Think about how much faster CI runs can be ;)
- You can avoid broken builds when someone deletes a commit, tag or repository

I'm going to explain how module servers work, introduce Athens, and show how to use it in action.

{{< screencast_bottom youtube_id="P3P9NINDW1k" github_episode="epsode28">}}
