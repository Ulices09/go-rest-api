package auth

import (
	"github.com/labstack/echo/v4"
)

func NewAuthRouter(e *echo.Echo, co AuthController) {
	postsGroup := e.Group("/auth")
	postsGroup.POST("/login", co.Login)
}
