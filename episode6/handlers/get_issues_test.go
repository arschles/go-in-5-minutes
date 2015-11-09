package handlers

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/arschles/testsrv"
	"github.com/google/go-github/github"
)

// TestAsciiCatTestSrv tests the AsciiCat handler by running the router in a real server.
// This test uses the arschles/testsrv (https://godoc.org/github.com/arschles/testsrv) library
// to run the server. That library uses net/http/httptest under the covers. See https://godoc.org/net/http/httptest#Server
//
// testsrv is useful for testing handlers that rely on complex router logic.
// If you use Gorilla Mux, this is a better option because the router does pre-processing such
// as extracting path variables for access by the handlers
func TestGetIssuesTestSrv(t *testing.T) {
	ghClient := github.NewClient(nil)
	// create the *entire* router here, which includes other routes
	r := NewRouter(ghClient)
	// start the server and have it run our router. because tests could run
	// concurrently, you should create and run one server per test case, and always
	// remember to close the server at the end of the test.
	// see https://godoc.org/github.com/arschles/testsrv for more documentation
	// on testsrv, and
	srv := testsrv.StartServer(r)
	// always close the server at the end of each test
	defer srv.Close()

	_, err := http.Get(fmt.Sprintf("%s/issues/arschles/go-in-5-minutes", srv.URLStr()))
	if err != nil {
		t.Fatalf("error executing GET on /issues/arschles/go-in-5-minutes [%s]", err)
	}
	reqs := srv.AcceptN(1, 1*time.Second)
	if len(reqs) < 1 {
		t.Fatal("expected one request to /issues/arschles/go-in-5-minutes, got none")
	}

	// do other validation on the request here
}
