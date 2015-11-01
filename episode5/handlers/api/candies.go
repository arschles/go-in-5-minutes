package api

import (
	"encoding/json"
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode5/models"
)

func Candies(db models.DB) http.Handler {
	type ret struct {
		Candies []string `json:"candies"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys, err := db.GetAllKeys()
		if err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
			return
		}
		// Note: this is a potentially large scale operation.
		// several improvements could be made:
		// - paginate the results, to provide an upper bound on amount of work in a single request
		// - send only the keys down to the browser, and have the browser do a GET on only the keys it needs
		candies := []string{}
		for _, key := range keys {
			candy := new(models.Candy)
			db.Get(key, candy)
			candies = append(candies, candy.Name)
		}
		if err := json.NewEncoder(w).Encode(ret{Candies: candies}); err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
		}
	})
}
