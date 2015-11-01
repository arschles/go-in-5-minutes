package api

import (
	"encoding/json"
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode5/models"
)

func CandyKeysList(db models.DB) http.Handler {
	type ret struct {
		Keys []string `json:"keys"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys, err := db.GetAllKeys()
		if err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
			return
		}
		if err := json.NewEncoder(w).Encode(ret{Keys: keys}); err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
		}
	})
}
