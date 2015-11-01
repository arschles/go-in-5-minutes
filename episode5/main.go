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
	"github.com/codegangsta/negroni"
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
	renderer := handlers.NewRenderRenderer("templates", []string{".html"}, handlers.Funcs, dev)
	var db models.DB
	if dev {
		db = models.NewInMemoryDB()
	} else {
		d, err := models.NewMongoDB(cfg.MongoURL)
		if err != nil {
			log.Fatalf("error connecting to Mongo [%s]", err)
			os.Exit(1)
		}
		db = d
	}

	r := mux.NewRouter()
	r.Handle("/", handlers.Index(renderer)).Methods("GET")
	r.Handle("/candies", handlers.Candies(renderer)).Methods("GET")

	r.Handle("/api/candies", api.Candies(db)).Methods("GET")
	r.Handle("/api/candy", api.CreateCandy(db)).Methods("PUT")

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, http.StatusNotFound, "not_found", map[string]string{
			"url": r.URL.String(),
		}, "layout")
	})

	n := negroni.Classic()
	n.UseHandler(r)
	hostStr := fmt.Sprintf(":%d", cfg.Port)
	n.Run(hostStr)
}
