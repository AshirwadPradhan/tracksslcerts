package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"

	"github.com/AshirwadPradhan/tracksslcerts/db"
	"github.com/AshirwadPradhan/tracksslcerts/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		return errors.New("template not found " + name)
	}
	// name of the base template in base.html
	return tmpl.ExecuteTemplate(w, "base", data)
}

func main() {

	s := db.NewSqliteUserStore()
	fmt.Println(s)

	e := echo.New()
	e.Static("/static", "static")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	templates := map[string]*template.Template{}
	templates["home"] = template.Must(template.ParseFiles("./templates/home.html", "./templates/base.html"))
	templates["create_account"] = template.Must(template.ParseFiles("./templates/base.html", "./templates/create_account.html"))
	templates["dashboard"] = template.Must(template.ParseFiles("./templates/base.html", "./templates/dashboard.html"))
	templates["user_login"] = template.Must(template.ParseFiles("./templates/base.html", "./templates/user_login.html"))

	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	e.GET("/", handlers.HomeHandler)
	e.GET("/dashboard", handlers.DashboardHandler)
	e.GET("/user-login", handlers.UserLoginHandler)
	e.GET("/create-account", handlers.CreateAccountHandler)
	e.POST("/register-user", handlers.RegisterUserHandler(s))

	e.Logger.Fatal(e.Start(":3000"))
}
