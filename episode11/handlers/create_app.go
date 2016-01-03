package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode11/db"
	"github.com/arschles/go-in-5-minutes/episode11/models"
	"github.com/gorilla/mux"
)

type CreateAppHandler struct {
	db db.DB
}

func NewCreateAppHandler(db db.DB) *CreateAppHandler {
	return &CreateAppHandler{db: db}
}

func (c *CreateAppHandler) RegisterRoute(r *mux.Router) {
	r.Handle("/apps", c).Methods("POST")
}

func (c *CreateAppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a := &models.App{}
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		http.Error(w, jsonErr(err, "couldn't decode JSON body"), http.StatusBadRequest)
		return
	}
	key := models.NewAppKey(a.Name)
	if err := c.db.Save(key, a); err != nil {
		http.Error(w, jsonErr(err, "couldn't save to the database"), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(jsonKVP("status", "created")))
}
