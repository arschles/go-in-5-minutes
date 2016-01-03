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

func main() {
	conf, err := GetConfig()
	if err != nil {
		log.Printf("Error getting config [%s]", err)
		os.Exit(1)
	}

	router := mux.NewRouter()

	redisOpts := &redis.Options{
		Addr:     fmt.Sprintf(conf.RedisHost),
		Password: conf.RedisPass,
		DB:       conf.RedisDB,
	}
	rawRedisClient := redis.NewClient(redisOpts)
	redisClient := db.NewRedis(rawRedisClient)

	cah := handlers.NewCreateAppHandler(redisClient)
	cah.RegisterRoute(router)

	gah := handlers.NewGetAppHandler(redisClient)
	gah.RegisterRoute(router)

	dah := handlers.NewDeleteAppHandler(redisClient)
	dah.RegisterRoute(router)

	portStr := fmt.Sprintf(":%d", conf.Port)
	log.Printf("Serving on %s", portStr)
	log.Fatal(http.ListenAndServe(portStr, router))
}
