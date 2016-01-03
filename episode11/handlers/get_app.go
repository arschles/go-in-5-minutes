package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode11/db"
	"github.com/arschles/go-in-5-minutes/episode11/models"
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
	name, ok := mux.Vars(r)[appNamePath]
	if !ok {
		http.Error(w, jsonErrStr("missing app name in path"), http.StatusBadRequest)
		return
	}
	app := &models.App{}
	if err := c.db.Get(models.NewAppKey(name), app); err != nil {
		http.Error(w, jsonErrStr(fmt.Sprintf("database error: %s", err)), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(app); err != nil {
		http.Error(w, jsonErr(err, "encoding json"), http.StatusInternalServerError)
		return
	}
}
