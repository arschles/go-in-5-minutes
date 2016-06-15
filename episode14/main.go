package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func hdl(t *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tplData := map[string]string{}
		// iterate through all query string values. overwrite keys that occur twice in the query string
		for keys, vals := range r.URL.Query() {
			for 
			tplData[k] = v
		}
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
