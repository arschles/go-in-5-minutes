package api

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/arschles/go-in-5-minutes/episode5/models"
)

func CreateCandy(db models.DB) http.Handler {
	type jsonInput struct {
		Name string `json:"name"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		in := jsonInput{}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			jsonErr(w, http.StatusBadRequest, err)
			return
		}
		candy := models.Candy{Name: in.Name}
		if _, err := db.Upsert(bson.NewObjectId().String(), candy); err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}
