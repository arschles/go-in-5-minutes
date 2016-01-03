package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arschles/go-in-5-minutes/episode11/db"
	"github.com/arschles/go-in-5-minutes/episode11/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/redis.v3"
)

// This server simulates an inventory system for applications inside a Platform as a Service (PaaS)

func main() {
	conf, err := GetConfig()
	if err != nil {
		log.Printf("Error getting config [%s]", err)
		os.Exit(1)
	}

	router := mux.NewRouter()

	var database db.DB
	switch conf.DBType {
	case "mem":
		database = db.NewMem()
	case "redis":
		redisOpts := &redis.Options{
			Addr:     fmt.Sprintf(conf.RedisHost),
			Password: conf.RedisPass,
			DB:       conf.RedisDB,
		}
		rawRedisClient := redis.NewClient(redisOpts)
		database = db.NewRedis(rawRedisClient)
	default:
		log.Printf("Error: no available DB type %s", conf.DBType)
		os.Exit(1)
	}

	cah := handlers.NewCreateAppHandler(database)
	cah.RegisterRoute(router)

	gah := handlers.NewGetAppHandler(database)
	gah.RegisterRoute(router)

	dah := handlers.NewDeleteAppHandler(database)
	dah.RegisterRoute(router)

	portStr := fmt.Sprintf(":%d", conf.Port)
	log.Printf("Serving on %s", portStr)
	log.Fatal(http.ListenAndServe(portStr, router))
}
