package httpapp

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const SessionToken = "session-token"

func GetSessionCookie(c echo.Context) (*http.Cookie, error) {
	return c.Cookie(SessionToken)
}

func SetSessionCookie(c echo.Context, value string) {
	c.SetCookie(&http.Cookie{
		Name:     SessionToken,
		Value:    value,
		Secure:   false, // TODO: poner true para producción
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})
}
