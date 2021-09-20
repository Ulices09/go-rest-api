package auth

import (
	httpapp "go-rest-api/internal/interface/http"

	"github.com/labstack/echo/v4"
)

func NewAuthRouter(
	e *echo.Echo,
	co AuthController,
	m httpapp.CustomMiddleware,
) {
	postsGroup := e.Group("/auth")
	postsGroup.POST("/login", co.Login)
	postsGroup.POST("/logout", co.Logout)
	postsGroup.POST("/me", co.Me, m.Auth())
}
