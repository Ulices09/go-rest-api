package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}
	httpMiddleware := InitMiddlware()
	e.Use(httpMiddleware.Logger())

	return e
}
