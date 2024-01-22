package handlers

import (
	"net/http"

	"github.com/AshirwadPradhan/tracksslcerts/auth"
	"github.com/AshirwadPradhan/tracksslcerts/db"
	"github.com/AshirwadPradhan/tracksslcerts/helpers"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func UserSignIn(db db.UserStorer) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: user sign in logic
		username := c.FormValue("username")
		password := c.FormValue("password")

		user, err := db.ReadUser(username)
		if err != nil {
			c.Logger().Error("cannot read user ", err)
			return c.String(http.StatusNotFound,
				helpers.FormMessageHTMXResponse("error", "Login Failed"))
		}
		if len(user.UserName) <= 0 {
			c.Logger().Error("username not found")
			return c.String(http.StatusNotFound,
				helpers.FormMessageHTMXResponse("error", "Login Failed"))
		}

		if checkpassword(password, user.HashedPassword) {
			c.Logger().Error("password does not match")
			return c.String(http.StatusBadRequest,
				helpers.FormMessageHTMXResponse("error", "Login Failed"))
		}

		if err = auth.GenerateTokenAndSetCookie(user, c); err != nil {
			c.Logger().Error("token is invalid")
			return c.String(http.StatusBadRequest,
				helpers.FormMessageHTMXResponse("error", "Login Failed"))
		}

		c.Response().Header().Set("HX-Redirect", "/dashboard")
		c.Response().Header().Set("HX-Location", "/dashboard")
		return c.Redirect(http.StatusOK, "/dashboard")
	}
}

func checkpassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
