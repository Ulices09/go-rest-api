package httpapp

import (
	"go-rest-api/internal/config"
	"go-rest-api/internal/core/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomMiddleware struct {
	config config.Config
}

func InitMiddlware(config config.Config) CustomMiddleware {
	return CustomMiddleware{config}
}

func (CustomMiddleware) Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}, uri=${uri}, status=${status}\n",
	})
}

func (m CustomMiddleware) CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     m.config.Host.AllowOrigins,
		AllowCredentials: true,
	})
}

func (CustomMiddleware) CSRF() echo.MiddlewareFunc {
	return middleware.CSRF()
}

func (CustomMiddleware) Recover() echo.MiddlewareFunc {
	return middleware.Recover()
}

func (CustomMiddleware) Secure() echo.MiddlewareFunc {
	return middleware.Secure()
}

func (m CustomMiddleware) Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenCookie, err := GetSessionCookie(c)

			if err != nil {
				return c.NoContent(http.StatusUnauthorized)
			}

			token, err := utils.VerifyJwt(tokenCookie.Value, m.config.Jwt.Secret)

			if err != nil {
				SetSessionCookie(c, "")
				return c.NoContent(http.StatusUnauthorized)
			}

			// TODO: ver si es necesario llamar a bdpara llenar el resto de los datos del usuario
			c.Set("user", token.Claims)
			return next(c)
		}
	}
}
