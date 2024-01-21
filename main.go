package main

import (
	"fmt"
	"net/http"

	"github.com/AshirwadPradhan/tracksslcerts/db"
	"github.com/AshirwadPradhan/tracksslcerts/handlers"
)

func main() {

	s := db.NewSqliteUserStore()
	fmt.Println(s)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/user-login", handlers.UserLoginHandler)
	http.HandleFunc("/dashboard", handlers.DashboardHandler)
	http.HandleFunc("/create-account", handlers.CreateAccountHandler)

	http.ListenAndServe(":3000", nil)
}
