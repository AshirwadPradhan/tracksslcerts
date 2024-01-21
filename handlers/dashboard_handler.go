package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", "dashboard.html")

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
	tmpl.ExecuteTemplate(w, "layout", nil)
}
