package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// hdl is a handler that calls t.Execute, passing all of the query string values as data to the template
func hdl(t *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Notice that t.Execute takes an io.Writer as its first argument. The variable w (an http.ResponseWriter) implements io.Writer, so we pass it and allow the implementation to call its Write function as necessary
		if err := t.Execute(w, r.URL.Query()); err != nil {
			http.Error(w, fmt.Sprintf("error executing template (%s)", err), http.StatusInternalServerError)
		}
	})
}

func main() {
	// This call does the following:
	// - template.ParseGlob parses *all* templates in the templates folder
	//		- Important if you have templates that depend on each other
	// - template.Must checks the error returned by ParseGlob and panics if it's non-nil. Otherwise returns the template that ParseGlob returned
	tpl := template.Must(template.New("site.html").ParseGlob("templates/*.html"))

	log.Printf("Server listening on port 8080")
	log.Printf("Example page: http://localhost:8080?a=b&a=c&c=d&d=e")
	if err := http.ListenAndServe(":8080", hdl(tpl)); err != nil {
		log.Fatalf("error running server (%s)", err)
	}
}
