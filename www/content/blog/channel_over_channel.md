+++
author = "Aaron Schlesinger"
date = "2016-09-05T10:11:58-07:00"
teaser = "Why passing a channel over a channel is a powerful concurrency pattern"
title = "Passing Channels over Channels"
type = "blog"

+++

# Passing Channels over Channels

[Channels](https://gobyexample.com/channels) in Go are one of the most powerful concurrency features of Go. They are essentially a shared, concurrency-safe queue (which I wrote about in my [last post](https://www.goin5minutes.com/blog/orthogonality/)) that you can use to pass data between your goroutines.

One of those pieces of "data" can be a channel itself, and that's what we're going to be talking about today.
