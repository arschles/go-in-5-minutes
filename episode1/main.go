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

	// http.ListenAndServe takes in an http.Handler as its second parameter.
	// since router is an http.ServeMux, and http.ServeMux is also a http.Handler
	// implementation (because it has a ServeHTTP func - https://godoc.org/net/http#ServeMux.ServeHTTP),
	// we can pass it here
	log.Printf("serving on port 8080")
	http.ListenAndServe(":8080", router)
}
