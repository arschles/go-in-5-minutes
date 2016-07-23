package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// hdl is a handler that calls t.Execute, passing all of the query string values as data to the template
func hdl(t *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tplData := map[string]string{}
		// iterate through all query string values. overwrite keys that occur twice in the query string
		for key, vals := range r.URL.Query() {
			for _, val := range vals {
				tplData[key] = val
			}
		}
		// Notice that t.Execute takes an io.Writer as its first argument. The variable w (an http.ResponseWriter) implements io.Writer, so we pass it and allow the implementation to call its Write function as necessary
		if err := t.Execute(w, tplData); err != nil {
			http.Error(w, fmt.Sprintf("error executing template (%s)", err), http.StatusInternalServerError)
		}
	})
}

func main() {
	tpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("error parsing templates (%s)", err)
	}

	log.Printf("Server listening on port 8080")
	http.ListenAndServe(":8080", hdl(tpl))
	tpl.Execute(os.Stdout, nil)
}
