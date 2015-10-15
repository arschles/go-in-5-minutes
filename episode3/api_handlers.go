package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getServer(w http.ResponseWriter, r *http.Request) {
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
