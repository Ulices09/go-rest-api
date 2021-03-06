package httpapp

import (
	"go-rest-api/internal/config"
	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/core/libs/jwt"
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
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			token, err := jwt.Verify(tokenCookie.Value, m.config.Jwt.Secret)

			if err != nil {
				SetSessionCookie(c, "")
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			// TODO: ver si es necesario llamar a bdpara llenar el resto de los datos del usuario
			c.Set("user", token.Claims)
			return next(c)
		}
	}
}

func (m CustomMiddleware) RoleGuard(roles ...interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			currentUser := entity.NewCurrentUser(c)
			hasRole := false

			for _, role := range roles {
				if role == currentUser.RoleId {
					hasRole = true
					break
				}
			}

			if !hasRole {
				return echo.NewHTTPError(http.StatusForbidden)
			}

			return next(c)
		}
	}
}
