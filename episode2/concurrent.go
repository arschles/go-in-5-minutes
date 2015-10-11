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

// getter executes the HTTP GET request against url, prints out the duration of the call
// (or an error message if it failed), and calls wg.Done() when it finishes.
// it's meant to be run in a goroutine.
func getter(url string, wg *sync.WaitGroup) {
	// this line tells godebug to replace it with a breakpoint when the code is instrumented
	_ = "breakpoint"
	defer wg.Done()
	start := time.Now()
	// actually do the GET
	if _, err := http.Get(url); err != nil {
		fmt.Printf("error GET-ing %s\n", url)
		return
	}
	fmt.Printf("duration to GET %s: %s\n", url, time.Now().Sub(start))
}

func main() {
	urls := []string{"https://google.com", "https://yahoo.com", "https://bing.com", "https://duckduckgo.com"}
	var wg sync.WaitGroup
	for _, url := range urls {
		// same as above, this line tells godebug to replace with a breakpoint when code is instrumented
		_ = "breakpoint"
		wg.Add(1)
		// like any program, the 'go' keyword tells the go runtime scheduler to run getter in the background.
		// godebug will not stop it from running the code concurrently, but since there's a breakpoint
		// in the getter func, we'll be able to inspect that function's context
		go getter(url, &wg)
	}
	wg.Wait() // block until all goroutines are done
	fmt.Println("done")
}
