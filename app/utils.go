package app

import (
	"go-rest-api/types/entity"

	"github.com/labstack/echo/v4"
)

func GetLoggedInUser(c echo.Context) entity.Claims {
	claims := c.Get("user").(*entity.Claims)
	return *claims
}
