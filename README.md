# Go In 5 Minutes

[![Build Status](https://travis-ci.org/arschles/go-in-5-minutes.svg?branch=master)](https://travis-ci.org/arschles/go-in-5-minutes) 
[![Build Status](https://cloud.drone.io/api/badges/arschles/go-in-5-minutes/status.svg)](https://cloud.drone.io/arschles/go-in-5-minutes)


Welcome Gophers! First off, lots of :heart: from me to you. I hope you enjoy
the screencasts, and as always, keep on rockin!

>Go in 5 Minutes was featured on [GoTime](https://gotime.fm). Check out the [episode](https://changelog.com/gotime/18)!


----

This repository has code and outlines for [Go In 5 Minutes Screencasts](htttp://bitly.com/goin5minutesyt). All content (code samples, outlines, etc...) is organized into folders, starting with [`episode0`](https://github.com/arschles/go-in-5-minutes/tree/master/episode0) and going from there.

Full descriptions and videos are at https://www.goin5minutes.com/screencasts/index.html.

Shortened URL for this repository: https://bitly.com/goin5minutes

Shortened URL for screencasts: https://bitly.com/goin5minutesyt

# Request A Screencast

If you've been interested in a Go package, best practice, or topic, I've probably had some experience with it, so I'd like
to hear from you. Especially if you'd like to see a screencast on it!

Please [submit an issue](https://github.com/arschles/go-in-5-minutes/issues) with a short description on the package/best practice/etc... that you've been thinking of.

Here are some example issues: https://github.com/arschles/go-in-5-minutes/labels/request%20a%20screencast.

**I'd love to hear from you!**

# Sponsors

Go in 5 Minutes has some folks who generously provide financial backing to help pay
for development costs. See [BACKERS.md](./BACKERS.md) for the list of those folks.

If you'd like to financially support the screencast, please consider buying my lessons:

- [The Ultimate Guide to Web Apps in Go](https://gum.co/hgHhj?wanted=true) - A subscription series that dives deep into the nuts and bolts of building full-stack web apps in Go. Since this is a subscription, you'll always be up to date on the latest in the fast moving world of Go web apps
- [The bundle of the first 10 episodes](https://gumroad.com/l/gifm-1-10?wanted=true) - A bundle of the first 10 episodes. All are DRM-free, HD, and you can download them and watch them anywhere. If you're just getting started with Go, I recommend this!

# Email Newsletter

I send out an email newsletter intermittently with screencasts, additional information and resources. I encourage you to sign up for it at https://www.goin5minutes.com/subscribe/index.html.

# Issues With Code, Documentation, etc...

If you see any problems with code, documentation, or anything else in this repository, please [submit an issue](https://github.com/arschles/go-in-5-minutes/issues) with the `bug` label and I'll fix it as soon as I can. Pull requests are also welcome.

# Build, Test and Run Instructions

All the folders that start with `episode` (such as [`episode0`](https://github.com/arschles/go-in-5-minutes/tree/master/episode0)) contain the outline and code samples for that episode, and all code samples can be built and run.

Unless otherwise specified in the `README.md` in the episode folder, the commands for building, testing and running simply use the `go` tool. All episodes are buildable and testable, but some don't have a `package main` because they're libraries, so they won't be runnable.

- build: `go build`
- test: `go test`
- run: `go build -o example && ./example`
