package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode11/db"
	"github.com/arschles/go-in-5-minutes/episode11/models"
	"github.com/gorilla/mux"
)

// GetAppHandler is the http.Handler that gets and returns existing applications
type GetAppHandler struct {
	db db.DB
}

// NewGetAppHandler initializes and returns a new GetAppHandler with the given database
func NewGetAppHandler(db db.DB) *GetAppHandler {
	return &GetAppHandler{db: db}
}

// RegisterRoute registers the appropriate route for this handler on the given router
func (c *GetAppHandler) RegisterRoute(r *mux.Router) {
	r.Handle(fmt.Sprintf("/app/{%s}", appNamePath), c).Methods("GET")
}

// ServeHTTP is the http.Handler interface implementation
func (c *GetAppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name, ok := mux.Vars(r)[appNamePath]
	if !ok {
		http.Error(w, jsonErrStr("missing app name in path"), http.StatusBadRequest)
		return
	}
	app := &models.App{}
	err := c.db.Get(models.NewAppKey(name), app)
	if err == db.ErrNotFound {
		http.Error(w, jsonErrStr(fmt.Sprintf("app %s not found", name)), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, jsonErrStr(fmt.Sprintf("database error: %s", err)), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(app); err != nil {
		http.Error(w, jsonErr(err, "encoding json"), http.StatusInternalServerError)
		return
	}
}
