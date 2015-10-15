package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// getServer is the handler to get a specific server by name. it's called the same way
// as if the standard net/http library called it, except we're guaranteed that it's called
// only if the specific request format (laid out in gorilla_server.go) is met
func getServer(w http.ResponseWriter, r *http.Request) {
	// mux.Vars gets a map of path variables by name. here "name" matches the {name} path
	// variable as seen in gorilla_server.go
	name, ok := mux.Vars(r)["name"]
	if !ok {
		http.Error(w, "name missing in URL path", http.StatusBadRequest)
		return
	}
	mx.RLock()
	defer mx.RUnlock()
	server, ok := servers[name]
	if !ok {
		http.Error(w, "no such server", http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(server); err != nil {
		log.Printf("[JSON Encoding Error] %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// reserveServer is the handler to reserve a specific server by name
func reserveServer(w http.ResponseWriter, r *http.Request) {
	name, ok := mux.Vars(r)["name"]
	if !ok {
		http.Error(w, "name missing in URL path", http.StatusBadRequest)
		return
	}
	mx.Lock()
	defer mx.Unlock()
	server, ok := servers[name]
	if !ok {
		http.Error(w, "no such server", http.StatusNotFound)
		return
	}
	server.Reserved = true
	server.NumReservations++
	if err := json.NewEncoder(w).Encode(server); err != nil {
		log.Printf("[JSON Encoding Error] %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// releaseServer is the handler to release a specific server by name
func releaseServer(w http.ResponseWriter, r *http.Request) {
	name, ok := mux.Vars(r)["name"]
	if !ok {
		http.Error(w, "name missing in URL path", http.StatusBadRequest)
		return
	}
	mx.Lock()
	defer mx.Unlock()
	server, ok := servers[name]
	if !ok {
		http.Error(w, "no such server", http.StatusNotFound)
		return
	}
	server.Reserved = false
	if err := json.NewEncoder(w).Encode(server); err != nil {
		log.Printf("[JSON Encoding Error] %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
