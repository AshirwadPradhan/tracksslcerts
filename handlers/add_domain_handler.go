package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleGetAddDomain(c echo.Context) error {
	return c.Render(http.StatusOK, "add_domain", nil)
}

func HandlePostAddDomain(c echo.Context) error {
	return c.Render(http.StatusOK, "add_domain", nil)
}
