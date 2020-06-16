package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	urlKey            = "url"
	defaultURLString  = "https://www.goin5minutes.com"
	timeoutKey        = "timeout"
	defaultTimeoutSec = 0
)

func getTimeout(def time.Duration, r *http.Request) time.Duration {
	toStr := r.URL.Query().Get(timeoutKey)
	toInt, err := strconv.Atoi(toStr)
	if err != nil {
		return def
	}
	return time.Duration(toInt) * time.Second
}

// client is a http.HandlerFunc that shows off how to use ctxhttp
func client(ctx context.Context, cl *http.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlStr := defaultURLString
		// get the timeout from the url query string
		timeout := getTimeout(defaultTimeoutSec, r)
		log.Printf("using timeout %s to request %s", timeout, urlStr)

		// create a new context for use with the http request
		// create the new context.Context with the timeout we fetched
		timeoutCtx, cancelFn := context.WithTimeout(ctx, timeout)
		// make sure to cancel the context when we're done. this is a good practice, EVEN IF IT HAS TIMED OUT
		defer cancelFn()

		// construct the http Request
		req, err := http.NewRequestWithContext(timeoutCtx, "GET", urlStr, nil)
		if err != nil {
			log.Printf("error creating request for %s (%s)", urlStr, err)
			http.Error(w, "error creating request", http.StatusInternalServerError)
			return
		}

		// execute the request. the request will be cancelled after timeout, even if the client's transport timeout has not been exceeded.
		resp, err := cl.Do(req)
		if err != nil {
			errStr := fmt.Sprintf("error making request for %s (%s)", urlStr, err)
			log.Println(errStr)
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// just copy the web page source to the client
		if _, err := io.Copy(w, resp.Body); err != nil {
			log.Printf("error copying (%s)", err)
			http.Error(w, "error copying", http.StatusInternalServerError)
			return
		}
	})
}
