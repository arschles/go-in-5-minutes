package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arschles/go-in-5-minutes/episode5/handlers"
	"github.com/arschles/go-in-5-minutes/episode5/handlers/api"
	"github.com/arschles/go-in-5-minutes/episode5/models"
	"github.com/arschles/go-in-5-minutes/episode5/util"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	cfg := util.Config{}
	if err := envconfig.Process("trickortreat", &cfg); err != nil {
		log.Fatalf("config error [%s]", err)
		os.Exit(1)
	}

	env, err := cfg.Env()
	if err != nil {
		log.Fatalf("config error [%s]", err)
		os.Exit(1)
	}
	dev := env == util.EnvDev
	renderer := handlers.NewRenderRenderer("templates", handlers.Funcs, dev)
	var db models.DB
	if dev {
		db = models.NewInMemoryDB()
	} else {
		db = models.NewMongoDB()
	}

	r := mux.NewRouter()
	r.Handle("/", handlers.Index(renderer, db))
	r.Handle("/api/candies", api.CandyList(db))

	hostStr := fmt.Sprintf(":%d", *port)
	log.Printf("serving on %s", hostStr)
	log.Fatal(http.ListenAndServe(hostStr, r))
}
