package auth

import (
	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/core/errors"
	httpapp "go-rest-api/internal/interface/http"
	"net/http"

	"github.com/labstack/echo/v4"
)

type controller struct {
	service AuthService
}

func NewAuthController(authService AuthService) AuthController {
	return &controller{service: authService}
}

func (co *controller) Login(c echo.Context) (err error) {
	data := new(LoginDto)

	if err = c.Bind(data); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	if err = c.Validate(data); err != nil {
		return
	}

	user, token, err := co.service.Login(data.Email, data.Password)

	if err != nil {
		return
	}

	httpapp.SetSessionCookie(c, token)
	return c.JSON(http.StatusOK, user)
}

func (co *controller) Logout(c echo.Context) (err error) {
	httpapp.SetSessionCookie(c, "")
	return c.NoContent(http.StatusOK)
}

func (co *controller) Me(c echo.Context) error {
	currentUser := entity.NewCurrentUser(c)
	user, err := co.service.Me(currentUser.Email)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
