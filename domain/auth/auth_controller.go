package auth

import (
	"go-rest-api/entity"
	"go-rest-api/utils"
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

	user, err := co.authService.Login(data.Email, data.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// TODO: mover a propio módulo para lógica de jwt y mover al service
	token, err := utils.SignAuthJwt(*user, "j7C6WjYm9DG9xWVe", 604800)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.SetCookie(&http.Cookie{
		Name:     "session-token",
		Value:    token,
		Secure:   false, // TODO: poner true para producción
		HttpOnly: true,
	})

	return c.JSON(http.StatusOK, user)
}

func (co *controller) Logout(c echo.Context) (err error) {
	c.SetCookie(&http.Cookie{
		Name:     "session-token",
		Secure:   false, // TODO: poner true para producción
		HttpOnly: true,
		MaxAge:   -1,
	})

	return c.NoContent(http.StatusOK)
}

func (co *controller) Test(c echo.Context) (err error) {
	claims := c.Get("user").(*entity.Claims)

	return c.String(http.StatusOK, "Your are logged in as: "+claims.Email)
}
