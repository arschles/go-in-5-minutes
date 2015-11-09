package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
)

func GetIssues(cl *github.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		// note that the checks on vars["org"] and vars["name"] are not strictly
		// necessary because Gorilla Mux will ensure that the values exist in the path,
		// we are doing the checks here as defensive programming (for example, if the
		// router is switched out later), and to ensure that any tests written directly
		// against this handler (as opposed to using net/http/httptest.Server or github.com/arschles/testsrv.Server)
		// will fail intelligently
		org, ok := vars["org"]
		if !ok {
			http.Error(w, "missing org in path", http.StatusBadRequest)
			return
		}
		name, ok := vars["name"]
		if !ok {
			http.Error(w, "missing repo name in path", http.StatusBadRequest)
			return
		}

		issues, _, err := cl.Issues.ListByRepo(org, name, nil)
		if err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
			return
		}
		titlesAndNumComments := make(map[string]int)
		for _, issue := range issues {
			titlesAndNumComments[*issue.Title] = *issue.Comments
		}
		if err := json.NewEncoder(w).Encode(titlesAndNumComments); err != nil {
			jsonErr(w, http.StatusInternalServerError, err)
		}
	})
}
