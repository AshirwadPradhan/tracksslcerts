package handlers

import (
	"net/http"

	"github.com/AshirwadPradhan/tracksslcerts/db"
	"github.com/AshirwadPradhan/tracksslcerts/types"
	"github.com/labstack/echo/v4"
)

func RegisterUserHandler(db db.UserStorer) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		email := c.FormValue("email")
		password := c.FormValue("password")

		if len(username) < 8 {
			c.Logger().Error("username length not satisfied")
			return c.String(http.StatusBadRequest,
				"<div id='form-message' class='text-red-500 text-sm mb-4'>Username not valid. It should be atleast 8 characters</div>")
		}

		u, err := db.ReadUser(username)
		if err != nil {
			c.Logger().Error("error in reading user ", err)
			return c.String(http.StatusBadRequest,
				"<div id='form-message' class='text-red-500 text-sm mb-4'>Error in creating user</div>")
		}
		if u.UserName == username {
			c.Logger().Error(" username already exists")
			return c.String(http.StatusBadRequest,
				"<div id='form-message' class='text-red-500 text-sm mb-4'>Username already exists</div>")
		}
		// TODO: hash the password and store

		newUser := types.NewUser(username, email, password)
		err = db.CreateUser(newUser)
		if err != nil {
			c.Logger().Error("error in creating user ", err)
			return c.String(http.StatusBadRequest,
				"<div id='form-message' class='text-red-500 text-sm mb-4'>Error in creating user</div>")
		}
		return c.String(http.StatusOK, "<div id='form-message' class='text-green-500 text-sm mb-4'>Account Created Successfully</div>")
	}
}
