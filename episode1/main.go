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
	// this creates a new Gorilla router. Use this router to build all the routes for your server,
	// similar to how a routes table would work in other systems.
	router := mux.NewRouter()
	// get the value of a key
	router.Handle("/{key}", handlers.GetKey(db)).Methods("GET").Schemes("http")
	// set the value of a key
	router.Handle("/{key}", handlers.PutKey(db)).Methods("PUT").Schemes("http")

	log.Printf("serving on port 8080")

	// http.ListenAndServe takes in an http.Handler as its second parameter.
	// since router is a Gorilla router (https://godoc.org/github.com/gorilla/mux#Router)
	// which is an http.Handler implementation, we can pass it here. Note that we could
	// also pass an http.ServeMux if we wanted to because it implements http.Handler
	// by providing a ServeHTTP func (https://godoc.org/net/http#ServeMux.ServeHTTP)
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
