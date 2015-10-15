package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// releaseAllServers is the handler to release the lock on all servers. it's an admin-only handler
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

// getAllServers is the handler to get the status of all servers. it's admin-only
func getAllServers(w http.ResponseWriter, r *http.Request) {
	mx.RLock()
	defer mx.RUnlock()
	if err := json.NewEncoder(w).Encode(servers); err != nil {
		log.Printf("[JSON Encoding Error] %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
