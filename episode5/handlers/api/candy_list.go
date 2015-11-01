package api

import (
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode5/models"
)

func CandyList(db models.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
