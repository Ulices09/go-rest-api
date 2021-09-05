package app

import (
	"fmt"
	"go-rest-api/config"
	"go-rest-api/entity"
	"net/http"

	"github.com/golang-jwt/jwt"
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
		AllowOrigins:     m.config.Host.AllowOrigins,
		AllowCredentials: true,
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

func (*CustomMiddleware) Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenCookie, err := c.Cookie("session-token")

			if err != nil {
				return c.NoContent(http.StatusUnauthorized)
			}

			// TODO: mover a propio módulo para lógica de jwt y mover al service
			token, err := jwt.ParseWithClaims(tokenCookie.Value, &entity.Claims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte("j7C6WjYm9DG9xWVe"), nil
			})

			if err != nil {
				return c.JSON(http.StatusUnauthorized, jwt.ErrInvalidKey)
			}

			// TODO: ver si es necesario llamar a bdpara llenar el resto de los datos del usuario
			c.Set("user", token.Claims)
			return next(c)
		}
	}
}
