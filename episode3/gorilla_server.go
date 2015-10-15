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
	// creates a new top level mux.Router. since a mux.Router implements the http.Handler interface,
	// we can pass it to http.ListenAndServe below
	router := mux.NewRouter()

	// configure the router to always run this handler when it couldn't match a request to any other handler
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("%s not found\n", r.URL)))
	})

	// create a subrouter just for standard API calls. subrouters are convenient ways to
	// group similar functionality together.
	apiRouter := router.Headers("Content-Type", "application/json").Subrouter()
	apiRouter.HandleFunc("/api/{name}", getServer).Methods("GET")
	apiRouter.HandleFunc("/api/{name}", reserveServer).Methods("POST")
	apiRouter.HandleFunc("/api/{name}", releaseServer).Methods("DELETE")

	adminAPIRouter := router.Headers("Content-Type", "application/json").MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		adminToken := r.Header.Get("X-ADMIN-TOKEN")
		if adminToken == "" {
			return false
		}
		// in a production setting, this could be a DB lookup to determine whether the token
		// is valid for admin usage
		return adminToken == "SuperAdmin"
	}).Subrouter()
	adminAPIRouter.HandleFunc("/api/admin/servers", getAllServers).Methods("GET")
	adminAPIRouter.HandleFunc("/api/admin/servers", releaseAllServers).Methods("DELETE")

	log.Printf("serving on port 8080")
	http.ListenAndServe(":8080", router)
}
