package main

import (
	"log"
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode11/db"
	"github.com/arschles/go-in-5-minutes/episode11/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	red := db.NewRedis()

	cah := handlers.NewCreateAppHandler(red)
	cah.RegisterRoute(router)

	gah := handlers.NewGetAppHandler(red)
	gah.RegisterRoute(router)

	dah := handlers.NewDeleteAppHandler(red)
	dah.RegisterRoute(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
