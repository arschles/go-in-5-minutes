package api

import (
	"encoding/json"
	"net/http"

	"labix.org/v2/mgo/bson"

	"github.com/arschles/go-in-5-minutes/episode5/models"
)

func CreateCandy(db models.DB) http.Handler {
	type jsonInput struct {
		Name string `json:"name"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := jsonInput{}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			// TODO: better error message
			jsonErr(w, http.StatusBadRequest, err)
			return
		}
		if _, err := db.Upsert(bson.NewObjectId().String(), models.Candy{Name: in.Name}); err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}
