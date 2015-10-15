package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// gorilla_server implements a REST API service that lets you reserve and release
// physical servers in a cloud computing system

type ServerStatus struct {
	NumReservations int  `json:"total_num_reservations"`
	Reserved        bool `json:"currently_reserved"`
}

// global variables are bad, using them here for brevity

// each server and the # times they've been reserved
var servers = map[string]*ServerStatus{
	"alice": &ServerStatus{},
	"bob":   &ServerStatus{},
	"carol": &ServerStatus{},
}

// the mutex to protect against concurrent access to servers
var mx sync.RWMutex

func main() {
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("%s not found\n", r.URL)))
	})

	router.HandleFunc("/api/{name}", getServer).Methods("GET")
	router.HandleFunc("/api/{name}", reserveServer).Methods("POST")
	router.HandleFunc("/api/{name}", releaseServer).Methods("DELETE")

	adminAPIRouter := router.Headers("X-ADMIN-TOKEN", "SuperAdmin").Subrouter()
	adminAPIRouter.HandleFunc("/api/admin/servers", getAllServers).Methods("GET")
	adminAPIRouter.HandleFunc("/api/admin/servers", releaseAllServers).Methods("DELETE")

	log.Printf("serving on port 8080")
	http.ListenAndServe(":8080", router)
}
