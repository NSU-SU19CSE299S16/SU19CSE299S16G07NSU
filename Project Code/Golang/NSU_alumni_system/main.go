package main

import (
	"html/template"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("templates/home.html", "templates/base.html"))

	tmpl.ExecuteTemplate(w, "base.html","")
}

func main() {
	http.HandleFunc("/", index_handler)

	http.ListenAndServe(":8000", nil)
}