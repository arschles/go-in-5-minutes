package handlers

import (
	"net/http"

	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
)

// NewRouter creates a new http.Handler that routes all valid HTTP paths to their respective
// handlers. It's best practice to register your handlers in a func like this so that you
// can test the resulting router (which itself is an http.Handler) directly.
//
// See this function used in the TestGetIssuesTestSrv unit test found in get_issues_test.go
func NewRouter(ghClient *github.Client) http.Handler {
	router := mux.NewRouter()
	router.Handle("/issues/{org}/{name}", GetIssues(ghClient))
	router.Handle("/ascii_cat", AsciiCat(ghClient))

	// handle other routes here

	return router
}
