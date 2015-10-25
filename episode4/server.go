package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"
)

func homeHandler(tpl *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r)
	})
}

func main() {
	flag.Parse()
	tpl := template.Must(template.ParseFiles("index.html"))
	h := newHub()
	router := http.NewServeMux()
	router.Handle("/", homeHandler(tpl))
	router.Handle("/ws", wsHandler{h: h})
	log.Printf("serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
