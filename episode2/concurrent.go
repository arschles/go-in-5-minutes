package main

// This program concurrently executes HTTP GET requests against various search
// engines and prints out the amount of time each GET took.
// We'll use it as an example to show godebug (https://github.com/mailgun/godebug) in action.

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type result struct {
	URL string
	Dur time.Duration
}

// getter executes the HTTP GET request against url, prints out the duration of the call
// (or an error message if it failed), and calls wg.Done() when it finishes.
// it's meant to be run in a goroutine.
func getter(url string, ch chan<- result) {
	defer close(ch)
	// this line tells godebug to replace it with a breakpoint when the code is instrumented
	start := time.Now()
	// actually do the GET
	if _, err := http.Get(url); err != nil {
		ch <- result{URL: url, Dur: time.Duration(-1)}
		return
	}
	ch <- result{URL: url, Dur: time.Now().Sub(start)}
}

func main() {
	urls := []string{"https://google.com", "https://yahoo.com", "https://bing.com", "https://duckduckgo.com"}
	durCh := make(chan result)
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		_ = "breakpoint"
		ch := make(chan result)
		go getter(url, ch)
		go func() {
			defer wg.Done()
			t := <-ch
			durCh <- t
		}()
	}
	go func() {
		wg.Wait()
		close(durCh)
	}()
	for t := range durCh {
		_ = "breakpoint"
		fmt.Println(t)
	}
}
