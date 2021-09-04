package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomMiddleware struct{}

// TODO: set config as parameter
func InitMiddlware() *CustomMiddleware {
	return &CustomMiddleware{}
}

func (*CustomMiddleware) Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}, uri=${uri}, status=${status}\n",
	})
}

func (*CustomMiddleware) CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.DefaultCORSConfig)
}

func (*CustomMiddleware) CSRF() echo.MiddlewareFunc {
	return middleware.CSRF()
}

func (*CustomMiddleware) Recover() echo.MiddlewareFunc {
	return middleware.Recover()
}

func (*CustomMiddleware) Secure() echo.MiddlewareFunc {
	return middleware.Secure()
}
