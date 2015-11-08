package main

import (
	"log"
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode6/handlers"
	"github.com/google/go-github/github"
)

func main() {
	ghClient := github.NewClient(nil)
	mux := http.NewServeMux()
	mux.Handle("/issues", handlers.GetIssues(ghClient))
	mux.Handle("/ascii_cat", handlers.AsciiCat(ghClient))
	mux.Handle("/meta", handlers.Meta(ghClient))

	log.Println("serving on port 8080")
	http.ListenAndServe(":8080", mux)
}
