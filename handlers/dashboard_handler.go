package handlers

import (
	"net/http"

	"github.com/AshirwadPradhan/tracksslcerts/db"
	"github.com/labstack/echo/v4"
)

func DashboardHandler(s db.DomainStorer) echo.HandlerFunc {

	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "dashboard", nil)
	}
}
