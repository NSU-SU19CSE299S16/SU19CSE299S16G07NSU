package templates_handler

import (
	chat_controller "NSU_Almuni/Controller"
	company_details "NSU_Almuni/Database_Handler"
	"fmt"
	"html/template"
	"net/http"
)

type Cnames struct {
	Names []string
	Details  map[string]interface{}
	Title string
	Condition bool
	Getname interface{}
}
var company_names = company_details.CompanyDetails()
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("templates/home.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "base.html","")
}
func CompanyHandler(w http.ResponseWriter, r *http.Request) {

	tmpl:= template.Must(template.ParseFiles("templates/auto_suggest_company_names.html"))
	data := Cnames{
		Names: nil,
		Details: company_names,
		Title:"SearchBar",
		Condition:false,
		Getname: nil,
	}
	for key,_ := range(company_names){
		data.Names = append(data.Names, key )
	}
	if r.Method != http.MethodPost {
		tmpl.Execute(w, data)
		return
	}
	names := r.FormValue("name")
	data.Getname = data.Details[names]
	fmt.Println(data.Getname)
	data.Condition = true
	tmpl.Execute(w, data)
}
func TemplatesHandler() {

	http.HandleFunc("/company_details",CompanyHandler)
	http.Handle("/messages", http.FileServer(http.Dir("./templates/chat")))
	http.HandleFunc("/new/user", chat_controller.RegisterNewUser)
	http.HandleFunc("/pusher/auth", chat_controller.PusherAuth)
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8000", nil)
}