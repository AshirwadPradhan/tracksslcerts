package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/dashboard.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}
