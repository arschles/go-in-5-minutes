package handlers

import (
	"fmt"
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode11/db"
	"github.com/gorilla/mux"
)

type DeleteAppHandler struct {
	db db.DB
}

func NewDeleteAppHandler(db db.DB) *DeleteAppHandler {
	return &DeleteAppHandler{db: db}
}

func (c *DeleteAppHandler) RegisterRoute(r *mux.Router) {
	r.Handle(fmt.Sprintf("/apps/{%s}", appNamePath), c).Methods("DELETE")
}

func (c *DeleteAppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
