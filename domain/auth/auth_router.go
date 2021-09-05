package auth

import (
	"go-rest-api/app"

	"github.com/labstack/echo/v4"
)

func NewAuthRouter(
	e *echo.Echo,
	co AuthController,
	m *app.CustomMiddleware,
) {
	postsGroup := e.Group("/auth")
	postsGroup.POST("/login", co.Login)
	postsGroup.POST("/logout", co.Logout)
	postsGroup.POST("/test", co.Test, m.Auth())
}
