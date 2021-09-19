package users

import (
	"github.com/labstack/echo/v4"
)

func NewUserRouter(e *echo.Echo, co UserController) {
	postsGroup := e.Group("/users")
	postsGroup.GET("", co.GetUsers)
	postsGroup.GET("/:id", co.GetUser)
	postsGroup.POST("", co.CreateUser)
}
