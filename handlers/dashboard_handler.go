package handlers

import (
	"net/http"

	"github.com/AshirwadPradhan/tracksslcerts/db"
	"github.com/labstack/echo/v4"
)

func DashboardHandler(s db.DomainStorer) echo.HandlerFunc {

	return func(c echo.Context) error {
		user, err := c.Cookie("ssl-cert-user")
		type data struct {
			Username string
		}

		if err != nil {
			return c.Render(http.StatusOK, "dashboard", data{Username: err.Error()})
		}
		return c.Render(http.StatusOK, "dashboard", data{Username: user.Value})
	}
}
