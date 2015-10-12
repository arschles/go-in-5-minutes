package handlers

import (
	"fmt"
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode1/storage"
)

// GetKey returns an http.Handler that can get a key registered by Gorilla mux
// as "key" in the path. It gets the value of the key from db
func GetKey(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing key name in query string", http.StatusBadRequest)
			return
		}
		val, err := db.Get(key)
		if err == storage.ErrNotFound {
			http.Error(w, "not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, fmt.Sprintf("error getting value from database: %s", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(val)
	})
}
