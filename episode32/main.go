package main

import (
	"log"
	"net/http"

	"golang.org/x/net/context"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/flush", flush)

	mux.Handle("/client", client(context.Background(), http.DefaultClient))

	log.Printf("serving on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
