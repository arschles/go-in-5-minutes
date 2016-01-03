package handlers

import (
	"fmt"
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode11/db"
	"github.com/gorilla/mux"
)

type GetAppHandler struct {
	db db.DB
}

func NewGetAppHandler(db db.DB) *GetAppHandler {
	return &GetAppHandler{db: db}
}

func (c *GetAppHandler) RegisterRoute(r *mux.Router) {
	r.Handle(fmt.Sprintf("/app/{%s}", appNamePath), c).Methods("GET")
}

func (c *GetAppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
