# Vendoring Dependencies in Go

[![Watch The Screencast](https://www.goin5minutes.com/img/watch-screencast.svg)](https://www.goin5minutes.com/screencast/episode_7_vendoring_your_dependencies_in_go/)

Go in 5 Minutes, episode 7.

This screencast shows how to vendor your dependencies in Go, the modern way.

>NOTE: This screencast is no longer relevant, with the advent of [Go Modules](https://golang.org/ref/mod). It is left here for historical reference only. Please use modules to handle (and vendor, if necessary) your dependencies.

Screencast video:
https://www.goin5minutes.com/screencast/episode_7_vendoring_your_dependencies_in_go/

## Email Newsletter

Quick reminder to sign up at https://www.goin5minutes.com/subscribe

The first email will go out this week!

## Outline

1. `go get`
2. Reproducible builds
3. A Rich History
  - [Godep](https://github.com/tools/godep) (my favorite) & Others
  - [Go 1.5 Vendor Experiment](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/edit)
  - [Glide](https://github.com/Masterminds/glide) - not the first or the last implementation
4. Example

## Note

Since dependency code is stored on Github, dependencies are mutable. I have checked in the vendor directory here so that you can compile & test this code. To experiment, just `rm -rf vendor` and run the `glide` commands.
