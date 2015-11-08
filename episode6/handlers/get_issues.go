package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/go-github/github"
)

func GetIssues(cl *github.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		issues, _, err := cl.Issues.ListByRepo("arschles", "go-in-5-minutes", nil)
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
