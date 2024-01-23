package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleGetCreateAccount(c echo.Context) error {
	return c.Render(http.StatusOK, "create_account", nil)
}
