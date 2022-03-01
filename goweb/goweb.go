package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>Hello, you've requested: %s\n</p>", r.URL.Path)
	fmt.Fprintf(w, "<h1>I hope you're happy!</h1>")
	fmt.Fprintf(w, "<p>Also %s add %s</p1>", "can", "<strong>variables and style</strong>")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	r.HandleFunc("/", index_handler)

	fs := http.FileServer(http.Dir("static/"))
	r.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8000", r)
}
