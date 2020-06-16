package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// flush is the http.HandlerFunc that shows off how to use an http.Flusher
func flush(w http.ResponseWriter, r *http.Request) {
	// tell curl to start accepting results
	w.WriteHeader(http.StatusOK)
	// tell curl to handle bodies as they are flushed
	w.Header().Set("Content-Type", "text/event-stream")
	// try to cast the ResponseWriter to a flusher
	flsh, ok := w.(http.Flusher)
	if !ok {
		// if w wasn't a flusher, bail out
		http.Error(w, "not a supported flusher!", http.StatusInternalServerError)
		return
	}
	// write and then flush!
	for i := 0; i < 10; i++ {
		toWrite := fmt.Sprintf("response %d\n", i)
		log.Printf("writing %s", toWrite)
		fmt.Fprintf(w, toWrite)
		flsh.Flush()
		time.Sleep(1 * time.Second)
	}
}
