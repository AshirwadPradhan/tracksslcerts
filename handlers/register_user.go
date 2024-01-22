package handlers

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/AshirwadPradhan/tracksslcerts/db"
	"github.com/AshirwadPradhan/tracksslcerts/types"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserHandler(db db.UserStorer) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		email := c.FormValue("email")
		password := c.FormValue("password")

		// check length of username
		if len(username) < 8 {
			c.Logger().Error("username length not satisfied")
			return c.String(http.StatusBadRequest,
				createAccountHTMXResponse("error", "Username should be atleast 8 charcters"))
		}

		u, err := db.ReadUser(username)
		if err != nil {
			c.Logger().Error("error in reading user ", err)
			return c.String(http.StatusBadRequest,
				createAccountHTMXResponse("error", "Error in creating user"))
		}
		if u.UserName == username {
			c.Logger().Error("username already exists")
			return c.String(http.StatusBadRequest,
				createAccountHTMXResponse("error", "Username already exists"))
		}

		if !validateEmail(email) {
			c.Logger().Error("email address not valid")
			return c.String(http.StatusBadRequest,
				createAccountHTMXResponse("error", "Email is not valid"))
		}
		// FIXME: different username can have same email id 
		// we are not checking for globally unique email id
		if u.Email == email {
			c.Logger().Error("email already exists")
			return c.String(http.StatusBadRequest,
				createAccountHTMXResponse("error", "Email is already registered"))
		}

		if !validatePassword(password) {
			c.Logger().Error("password validation failed")
			return c.String(http.StatusBadRequest,
				createAccountHTMXResponse("error", "Invalid Password"))
		}

		hpass, err := hashPassword(password)
		if err != nil {
			c.Logger().Error("error in hashing password")
			return c.String(http.StatusBadRequest,
				createAccountHTMXResponse("error", "Invalid Password"))
		}

		newUser := types.NewUser(username, email, hpass)
		err = db.CreateUser(newUser)
		if err != nil {
			c.Logger().Error("error in creating user ", err)
			return c.String(http.StatusBadRequest,
				createAccountHTMXResponse("error", "Error in creating user"))
		}

		c.Logger().Info("account created successfully")
		return c.String(http.StatusOK, createAccountHTMXResponse("ok", "Account created successfully"))
	}
}

func createAccountHTMXResponse(kind string, message string) string {
	if kind == "ok" {
		return fmt.Sprintf("<div id='form-message' class='text-green-500 text-sm mb-4'>%s</div>", message)
	} else {
		return fmt.Sprintf("<div id='form-message' class='text-red-500 text-sm mb-4'>%s</div>", message)
	}
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validatePassword(password string) bool {
	if len(password) > 7 && len(password) < 69 {
		return true
	}
	return false
}

func hashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(b), err
}

func checkpassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
