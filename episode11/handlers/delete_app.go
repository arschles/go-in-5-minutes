package handlers

import (
	"fmt"
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode11/db"
	"github.com/arschles/go-in-5-minutes/episode11/models"
	"github.com/gorilla/mux"
)

type DeleteAppHandler struct {
	db db.DB
}

func NewDeleteAppHandler(db db.DB) *DeleteAppHandler {
	return &DeleteAppHandler{db: db}
}

func (c *DeleteAppHandler) RegisterRoute(r *mux.Router) {
	r.Handle(fmt.Sprintf("/app/{%s}", appNamePath), c).Methods("DELETE")
}

func (c *DeleteAppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name, ok := mux.Vars(r)[appNamePath]
	if !ok {
		http.Error(w, jsonErrStr("app name not found in the path"), http.StatusBadRequest)
		return
	}
	if err := c.db.Delete(models.NewAppKey(name)); err != nil {
		http.Error(w, jsonErr(err, "DB error when trying to delete"), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
