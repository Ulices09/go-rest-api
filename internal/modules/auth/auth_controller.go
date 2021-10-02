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

// Login godoc
// @Summary Login
// @Tags auth
// @Param default body auth.LoginRequest true "data"
// @Success 200
// @Failure default {object} errors.CustomError
// @Router /auth/login [post]
func (co *controller) Login(c echo.Context) (err error) {
	data := new(LoginRequest)

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

// Logout godoc
// @Summary Logout
// @Tags auth
// @Success 200
// @Failure default {object} errors.CustomError
// @Router /auth/logout [post]
func (co *controller) Logout(c echo.Context) (err error) {
	httpapp.SetSessionCookie(c, "")
	return c.NoContent(http.StatusOK)
}

// Me godoc
// @Summary Get current user
// @Tags auth
// @Security auth-token
// @Success 200 {object} entity.User
// @Failure default {object} errors.CustomError
// @Router /auth/me [get]
func (co *controller) Me(c echo.Context) error {
	currentUser := entity.NewCurrentUser(c)
	user, err := co.service.Me(currentUser.Email)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
