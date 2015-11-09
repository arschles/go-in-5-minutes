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
