package auth

import (
	httpapp "go-rest-api/internal/interface/http"
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

	httpapp.SetSessionCookie(c, token)
	return c.JSON(http.StatusOK, user)
}

func (co *controller) Logout(c echo.Context) (err error) {
	httpapp.SetSessionCookie(c, "")
	return c.NoContent(http.StatusOK)
}

func (co *controller) Me(c echo.Context) error {
	userClaims := httpapp.GetLoggedInUser(c)
	user, err := co.authService.Me(userClaims.Email)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}