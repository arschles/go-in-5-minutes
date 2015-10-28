package main

// Serve is a very simple static file server in go
// Usage:
// -p="8100": port to serve on
// -d=".":    the directory of static files to host
// Navigating to http://localhost:8100 will display the index.html or directory
// listing file.
// thanks to https://gist.github.com/paulmach/7271283 for this

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8080", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
