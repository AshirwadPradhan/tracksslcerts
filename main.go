package main

import (
	"net/http"

	"github.com/AshirwadPradhan/tracksslcerts/handlers"
)

func main() {

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/userlogin", handlers.UserLoginHandler)
	http.HandleFunc("/dashboard", handlers.DashboardHandler)

	http.ListenAndServe(":3000", nil)
}
