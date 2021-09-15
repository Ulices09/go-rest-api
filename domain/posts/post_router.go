package posts

import (
	"go-rest-api/app"

	"github.com/labstack/echo/v4"
)

func NewPostRouter(
	e *echo.Echo,
	co PostController,
	m app.CustomMiddleware,
) {
	postsGroup := e.Group("/posts")
	postsGroup.GET("", co.GetPosts)
	postsGroup.GET("/:id", co.GetPost)
	postsGroup.POST("", co.CreatePost, m.Auth())
}
