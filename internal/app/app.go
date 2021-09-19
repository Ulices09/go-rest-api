package app

import (
	"go-rest-api/internal/config"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func New(config config.Config, cMiddleware CustomMiddleware) *echo.Echo {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(cMiddleware.Logger())
	e.Use(cMiddleware.CORS())
	e.Use(cMiddleware.Recover())
	e.Use(cMiddleware.Secure())

	return e
}
