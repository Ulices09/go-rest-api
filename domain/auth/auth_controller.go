package auth

import (
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

	users, err := co.authService.Login(data.Email, data.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}
