package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-github/github"
)

// TestAsciiCatRespRecorder uses net/http/httptest ResponseRecorder
// (https://godoc.org/net/http/httptest#ResponseRecorder) to test the AsciiCat
// handler directly.
//
// ResponseRecorder is useful for direct testing of handlers,
// but doesn't provide a complete solution when the router itself handles complex logic.
// See TestGetIssuesTestSrv in get_issues_test.go for an example of testing complex router logic
func TestAsciiCatRespRecorder(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/ascci_cat", nil)
	if err != nil {
		t.Fatalf("error constructing test HTTP request [%s]", err)
	}
	ghClient := github.NewClient(nil)
	AsciiCat(ghClient).ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("expected code %d, got %d", http.StatusOK, w.Code)
	}
	bodyStr := string(w.Body.Bytes())
	if len(bodyStr) <= 0 {
		t.Fatalf("expected non-empty response body")
	}
	expectedCat, _, err := ghClient.Octocat("Hello, Go In 5 Minutes Viewer!")
	if err != nil {
		t.Fatalf("error getting expected octocat string [%s]", err)
	}
	if bodyStr != expectedCat {
		t.Fatalf("got unexpected octocat string [%s]", bodyStr)
	}
}
