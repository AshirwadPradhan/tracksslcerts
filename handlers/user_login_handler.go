package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleGetUserLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "user_login", nil)
}
