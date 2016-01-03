package handlers

import (
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode11/db"
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

}
