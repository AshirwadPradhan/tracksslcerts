package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleGetHome(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}
