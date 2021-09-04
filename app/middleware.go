package app

import (
	"go-rest-api/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomMiddleware struct {
	config config.Config
}

func InitMiddlware(config config.Config) *CustomMiddleware {
	return &CustomMiddleware{config}
}

func (*CustomMiddleware) Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}, uri=${uri}, status=${status}\n",
	})
}

func (m *CustomMiddleware) CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: m.config.Host.AllowOrigins,
	})
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
