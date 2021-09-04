package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}
	cMiddleware := InitMiddlware()
	e.Use(cMiddleware.Logger())
	e.Use(cMiddleware.CORS())
	e.Use(cMiddleware.CSRF())
	e.Use(cMiddleware.Recover())
	e.Use(cMiddleware.Secure())

	return e
}
