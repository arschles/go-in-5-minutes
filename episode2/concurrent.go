package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func getter(url string, wg sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	if _, err := http.Get(url); err != nil {
		fmt.Printf("error GET-ing %s\n", url)
		return
	}
	fmt.Printf("duration to GET %s: %s\n", url, time.Now().Sub(start))
}

func main() {
	urls := []string{"https://google.com", "https://yahoo.com", "https://bing.com", "https://duckduckgo.com"}
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go getter(url, wg)
	}
	wg.Wait()
}
