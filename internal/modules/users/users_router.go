package users

import (
	"go-rest-api/internal/core/constants"
	httpapp "go-rest-api/internal/interface/http"

	"github.com/labstack/echo/v4"
)

func NewUserRouter(
	e *echo.Echo,
	co UserController,
	m httpapp.CustomMiddleware,
) {
	postsGroup := e.Group("/users")
	postsGroup.GET("", co.GetUsers)
	postsGroup.GET("/:id", co.GetUser)
	postsGroup.POST(
		"",
		co.CreateUser,
		m.Auth(),
		m.RoleGuard(constants.RoleAdmin),
	)
}
