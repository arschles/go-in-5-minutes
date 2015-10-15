package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func releaseAllServers(w http.ResponseWriter, r *http.Request) {
	mx.Lock()
	defer mx.Unlock()
	for _, status := range servers {
		status.Reserved = false
	}
	if err := json.NewEncoder(w).Encode(servers); err != nil {
		log.Printf("[JSON Encoding Error] %s", err)
		http.Error(w, err.Error(), http.StatusNoContent)
	}
}

func getAllServers(w http.ResponseWriter, r *http.Request) {
	mx.RLock()
	defer mx.RUnlock()
	if err := json.NewEncoder(w).Encode(servers); err != nil {
		log.Printf("[JSON Encoding Error] %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
