package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/user_login.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}
