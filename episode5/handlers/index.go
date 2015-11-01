package handlers

import (
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode5/models"
)

func Index(ren Renderer, db models.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
