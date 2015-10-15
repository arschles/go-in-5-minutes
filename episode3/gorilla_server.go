package main

import (
	"encoding/json"
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

func reserveServer(w http.ResponseWriter, r *http.Request) {
	name, ok := mux.Vars(r)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	mx.Lock()
	defer mx.Unlock()

}

func releaseServer(w http.ResponseWriter, r *http.Request) {
	name, ok := mux.Vars(r)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	mx.Lock()
	defer mx.Unlock()
}

func releaseAllServers(w http.ResponseWriter, r *http.Request) {
	mx.Lock()
	defer mx.Unlock()
	for _, status := range servers {
		status.Reserved = false
	}
	w.Write(http.StatusNoContent)
}

func getAllServers(w http.ResponseWriter, r *http.Request) {
	mx.RLock()
	defer mx.RUnlock()
	if err := json.Encoder(w).Encode(servers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	const adminToken = "SuperAdmin"

	router := mux.NewRouter()
	router.HandleFunc("/server/{name}", reserveServer).Methods("POST")
	router.HandleFunc("/server/{name}", releaseServer).Methods("DELETE")
	router.HandleFunc("/servers", getAllServers).Methods("GET").Headers("X-ADMIN-TOKEN", adminToken)
	router.HandleFunc("/servers", releaseAllServers).Methods("DELETE").Headers("X-ADMIN-TOKEN", adminToken)
	http.ListenAndServe(":8080", router)
}
