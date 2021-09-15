package auth

import (
	"go-rest-api/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

type controller struct {
	authService AuthService
}

func NewAuthController(authService AuthService) AuthController {
	return &controller{authService}
}

func (co *controller) Login(c echo.Context) (err error) {
	data := new(LoginDto)

	if err = c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(data); err != nil {
		return err
	}

	user, token, err := co.authService.Login(data.Email, data.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.SetCookie(&http.Cookie{
		Name:     "session-token",
		Value:    token,
		Secure:   false, // TODO: poner true para producción
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})

	return c.JSON(http.StatusOK, user)
}

func (co *controller) Logout(c echo.Context) (err error) {
	c.SetCookie(&http.Cookie{
		Name:     "session-token",
		Secure:   false, // TODO: poner true para producción
		HttpOnly: true,
		MaxAge:   -1,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})

	return c.NoContent(http.StatusOK)
}

func (co *controller) Me(c echo.Context) error {
	claims := c.Get("user").(*entity.Claims)
	user, err := co.authService.Me(claims.Email)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
