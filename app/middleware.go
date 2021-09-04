package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpMiddleware struct{}

// TODO: set config as parameter
func InitMiddlware() *HttpMiddleware {
	return &HttpMiddleware{}
}

func (*HttpMiddleware) Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}, uri=${uri}, status=${status}\n",
	})
}
