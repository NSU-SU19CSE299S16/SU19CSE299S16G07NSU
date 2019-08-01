package main

import (
	company_details "NSU_Almuni/Database_Handler"
	"html/template"
	"net/http"
)

type Cnames struct {
	Details []string
	Title string
}
var company_names = company_details.CompanyDetails()
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("templates/home.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "base.html","")
}
func CompanyHandler(w http.ResponseWriter, r *http.Request) {
	//var tmpl *template.Template
	data := Cnames{
		Details: company_names,
		Title:"SearchBar",
	}
	tmpl := template.Must(template.ParseFiles("templates/auto_suggest_company_names.html"))
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/company_details",CompanyHandler)

	http.ListenAndServe(":8000", nil)
}