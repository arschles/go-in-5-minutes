package main

import (
	"log"
	"net/http"
	"os"

	"github.com/arschles/go-in-5-minutes/episode11/db"
	"github.com/arschles/go-in-5-minutes/episode11/handlers"
	"github.com/gorilla/mux"
)

func main() {
	conf, err := GetConfig()
	if err != nil {
		log.Printf("Error getting config [%s]", err)
		os.Exit(1)
	}

	router := mux.NewRouter()

	red := db.NewRedis()

	cah := handlers.NewCreateAppHandler(red)
	cah.RegisterRoute(router)

	gah := handlers.NewGetAppHandler(red)
	gah.RegisterRoute(router)

	dah := handlers.NewDeleteAppHandler(red)
	dah.RegisterRoute(router)

	portStr := fmt.Sprintf(":%d", conf.Port)
	log.Printf("Serving on %s", portStr)
	log.Fatal(http.ListenAndServe(portStr, router))
}
