+++
author = "Aaron Schlesinger"
date = "2016-09-05T10:11:58-07:00"
teaser = "Why passing a channel over a channel is a powerful concurrency pattern"
title = "Passing Channels over Channels"
type = "blog"

+++

# Passing Channels over Channels

As most know, [channels](https://gobyexample.com/channels) are one of the most powerful concurrency features in Go. Armed with Goroutines and the `select` statement, you can build correct, efficient and understandable concurrent programs that do complex things.

In essence, a channel is a shared, concurrency-safe queue. Its primary purpose is to pass data across concurrency boundaries (i.e. between goroutines). Another way to say that is: you can send or receive an instance of any `type` on a channel. I'm going to focus on sending that `chan` type over a channel.

# Why

One simple reason you'd send a `chan` on a `chan` is to tell a goroutine to do work and then get an acknowledgement (ack hereafter) that it's finished doing that work.

Here's what such a channel looks like in Go code:

```go
chanOverChan := make(chan chan int)
```

In English, this code means: "a channel on which you can send or receive a channel of ints". When you see code that looks like the above, it's a safe bet that the sender is telling the receiver to do some computation and send the results to another goroutine, which may be the sender. We're going to focus on case where the sender is the receiver that the ack is forwarded to.

# Patterns

You won't always see a simple `chan chan int`. Sometimes, the ack channel is stored inside a struct:

```go
type data struct {
  retCh chan<- int
}
dataCh := make(chan data)
```

And you might see the channel completely abstracted by a `func`:

```go
type abstractedCh := chan func(int)
```

In this case, the sender can capture the channel inside the `func(int)` if they want -- or they can send any other implementation they want. This strategy is called a [function closure](https://en.wikipedia.org/wiki/Closure_(computer_programming)), and is extremely flexible.

# In Action

Below are some code examples using the 3 strategies. In each case, We'll simulate the work using a simple `time.Sleep`.

## Style 1: Using a Channel Inside a Channel

Here's the simplest of the patterns in action. Generally this style will be easiest to read and understand, but it has some limits:

- Each `doStuff` goroutine sleeps for a set amount of time. You can't change the sleep time when you send on `ch`
- Each `doStuff` goroutine can _only_ receive a `chan time.Duration` -- no more data than that. We'll address that problem in the next style.

```go
package main

import (
	"log"
	"sync"
	"time"
)

// the function to be run inside a goroutine. It receives a channel on ch, sleeps for t, then sends t on the channel it received
func doStuff(t time.Duration, ch <-chan chan time.Duration) {
	ac := <-ch
	time.Sleep(t)
	ac <- t
}

func main() {
	// create the channel-over-channel type
	sendCh := make(chan chan time.Duration)

	// start up 10 doStuff goroutines
	for i := 0; i < 10; i++ {
		go doStuff(time.Duration(i+1)*time.Second, sendCh)
	}

	// send channels to each doStuff goroutine. doStuff will "ack" by sending its sleep time back
	recvCh := make(chan time.Duration)
	for i := 0; i < 10; i++ {
		sendCh <- recvCh
	}

	// receive on each channel we previously sent. this is where we receive the ack that doStuff sent back above
	var wg sync.WaitGroup // use this to block until all goroutines have received the ack and logged
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dur := <-recvCh
			log.Printf("slept for %s", dur)
		}()
	}
	wg.Wait()
}
```

See this code in action at https://play.golang.org/p/-1lY-4gd4N.

## Style 2: Using a Channel Stored Inside a Struct

This code will look almost identical to the previous snippet, with 2 exceptions:

- The ack channel will be stored inside a `struct`
- The sleep time will be stored inside that same `struct`, so we can pass it over the `channel`
  - This makes the code more flexible, because we can tell `doStuff` how long to sleep when we _send_ to it, rather than when we start it

```go
package main

import (
	"log"
	"sync"
	"time"
)

// the struct that we'll pass over a channel to a goroutine running doStuff
type process struct {
	dur time.Duration
	ch  chan time.Duration
}

// the goroutine function. will receive a process struct 'p' on ch, sleep for p.dur, then send p.dur on p.ch
func doStuff(ch <-chan process) {
	proc := <-ch
	time.Sleep(proc.dur)
	proc.ch <- proc.dur
}

func main() {
	// start up the goroutines
	sendCh := make(chan process)
	for i := 0; i < 10; i++ {
		go doStuff(sendCh)
	}

	// store an array of each struct we sent to the goroutines
	processes := make([]process, 10)
	for i := 0; i < 10; i++ {
		dur := time.Duration(i+1) * time.Second
		proc := process{dur: dur, ch: make(chan time.Duration)}
		processes[i] = proc
		sendCh <- proc
	}

	// recieve on each struct's ack channel
	var wg sync.WaitGroup // use this to block until all goroutines have received the ack and logged
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(ch <-chan time.Duration) {
			defer wg.Done()
			dur := <-ch
			log.Printf("slept for %s", dur)
		}(processes[i].ch)
	}
	wg.Wait()
}
```

See this code in action at https://play.golang.org/p/bJoiGP9ua2.

## Style 3: Using a Channel Inside a Function Closure

This code will look different from the previous examples, because the `doStuff` function won't know _anything_ about a return channel. That fact is both good and bad. On the up side, you can change your code later to do anything you want inside that function (e.g. good for testing!), but on the down side, you can't pass dynamic `time.Duration`s into the `doStuff` goroutines, as you could in the previous example.

```go
package main

import (
	"log"
	"sync"
	"time"
)

func doStuff(dur time.Duration, ch <-chan func(time.Duration)) {
	ackFn := <-ch
	time.Sleep(dur)
	ackFn(dur)
}

func main() {
	// start up the doStuff goroutines
	sendCh := make(chan func(time.Duration))
	for i := 0; i < 10; i++ {
		dur := time.Duration(i+1) * time.Second
		go doStuff(dur, sendCh)
	}

	// create the channels that will be closed over, create functions that close over each channel, then send them to the doStuff goroutines
	recvChs := make([]chan time.Duration, 10)
	for i := 0; i < 10; i++ {
		recvCh := make(chan time.Duration)
		recvChs[i] = recvCh
		fn := func(dur time.Duration) {
			recvCh <- dur
		}
		sendCh <- fn
	}

	// receive on the closed-over functions
	var wg sync.WaitGroup // use this to block until all goroutines have received the ack and logged
	for _, recvCh := range recvChs {
		wg.Add(1)
		go func(recvCh <-chan time.Duration) {
			defer wg.Done()
			dur := <-recvCh
			log.Printf("slept for %s", dur)
		}(recvCh)
	}
	wg.Wait()
}
```

See this code in action at https://play.golang.org/p/JAtGxdBVRW.

# Summary

There are uses for this channel-over-channel strategy, but the ack one is simple and powerful. Further, in many cases when you need to "return" something to another goroutine, sending it a `chan` on which it can return a value is often the easiest way to do it. This pattern can even be useful when you want to wait for a goroutine to ack its completion. Note, however, that you can also do ack-ing with a `sync.WaitGroup`.
