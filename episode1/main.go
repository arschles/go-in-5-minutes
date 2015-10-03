package main

import (
	"log"
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode1/handlers"
	"github.com/arschles/go-in-5-minutes/episode1/storage"
	"github.com/gorilla/mux"
)

func main() {
	// this creates the backend storage system
	db := storage.NewInMemoryDB()
	// this creates a new http.ServeMux, but it provides convenience functionality
	// for building a proper RESTful HTTP server
	router := mux.NewRouter()
	router.Handle("/{key}", handlers.GetKey(db)).Methods("GET") // get the value of a key
	router.Handle("/{key}", handlers.PutKey(db)).Methods("PUT") // set the value of a key

	log.Printf("serving on port 8080")

	// http.ListenAndServe takes in an http.Handler as its second parameter.
	// since router is a Gorilla router (https://godoc.org/github.com/gorilla/mux#Router)
	// which is an http.Handler implementation, we can pass it here. Note that we could
	// also pass an http.ServeMux if we wanted to because it implements http.Handler
	// by providing a ServeHTTP func (https://godoc.org/net/http#ServeMux.ServeHTTP)
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
