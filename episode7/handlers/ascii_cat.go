package handlers

import (
	"net/http"

	"github.com/google/go-github/github"
)

func AsciiCat(cl *github.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello, Go In 5 Minutes Viewer!"
		cat, _, err := cl.Octocat(msg)
		if err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cat))
	})
}
