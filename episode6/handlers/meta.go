package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/go-github/github"
)

func Meta(cl *github.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		meta, _, err := cl.APIMeta()
		if err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
			return
		}
		if err := json.NewEncoder(w).Encode(meta); err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
		}
	})
}
