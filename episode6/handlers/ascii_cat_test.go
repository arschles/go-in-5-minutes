package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-github/github"
)

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
