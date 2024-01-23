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
		username := c.FormValue("username")
		password := c.FormValue("password")

		user, err := db.ReadUser(username)
		if err != nil {
			c.Logger().Error("cannot read user ", err)
			return errOnValidation(c)
		}
		if len(user.UserName) <= 0 {
			c.Logger().Error("username not found")
			return errOnValidation(c)
		}

		if checkpassword(password, user.HashedPassword) {
			c.Logger().Error("password does not match")
			return errOnValidation(c)
		}

		if err = auth.GenerateTokenAndSetCookie(user, c); err != nil {
			c.Logger().Error("token is invalid")
			return errOnValidation(c)
		}

		// We need to redirect here to /dashboard on successful verification
		// Since we are using htmx, we need to set HX-Redirect header so that
		// htmx can redirect to correct page on form successful submission
		// We cannot user c.Redirect() as htmx only redirects on 2xx response
		// and c.Redirect() expects 3xx response code.
		// Therefore we will use htmx for redirect by setting HX-Redirect header
		// and echo will return no-content
		c.Response().Header().Set("HX-Redirect", "/dashboard")
		c.Response().Header().Set("HX-Location", "/dashboard")
		return c.NoContent(http.StatusOK)
	}
}

func errOnValidation(c echo.Context) error {
	return c.String(http.StatusBadRequest,
		helpers.FormMessageHTMXResponse("error", "Username or Password incorrect"))
}

func checkpassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
