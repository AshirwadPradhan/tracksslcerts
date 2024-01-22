package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/AshirwadPradhan/tracksslcerts/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	accessTokenCookieName = "ssl-cert-at"
	userCookieName        = "ssl-cert-user"
)

var jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

type Claims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func GetJWTSecret() string {
	return jwtSecretKey
}

func GenerateTokenAndSetCookie(user *types.User, c echo.Context) error {
	accessToken, exp, err := generateAccessToken(user, c)
	if err != nil {
		c.Logger().Error("error in generating token ", err)
	}

	setTokenCookie(accessTokenCookieName, accessToken, exp, c)
	setUserCookie(userCookieName, user, exp, c)

	return nil
}

func JWTErrorChecker(c echo.Context, err error) error {
	return c.Redirect(http.StatusMovedPermanently, c.Echo().Reverse("userLoginForm"))
}

func generateAccessToken(user *types.User, c echo.Context) (string, time.Time, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	return generateToken(user, expirationTime, []byte(GetJWTSecret()), c)
}

func generateToken(user *types.User, expirationTime time.Time, secret []byte, c echo.Context) (string, time.Time, error) {
	claims := &Claims{
		Name: user.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenStr, expirationTime, nil
}

func setTokenCookie(accessTokenCookieName string, accessToken string, exp time.Time, c echo.Context) {
	cookie := http.Cookie{}
	cookie.Name = accessTokenCookieName
	cookie.Value = accessToken
	cookie.Expires = exp
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(&cookie)
}

func setUserCookie(userCookieName string, user *types.User, exp time.Time, c echo.Context) {
	cookie := http.Cookie{}
	cookie.Name = userCookieName
	cookie.Value = user.UserName
	cookie.Expires = exp
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(&cookie)
}
