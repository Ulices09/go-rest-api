package posts

import (
	httpapp "go-rest-api/internal/http"

	"github.com/labstack/echo/v4"
)

func NewPostRouter(
	e *echo.Echo,
	co PostController,
	m httpapp.CustomMiddleware,
) {
	postsGroup := e.Group("/posts")
	postsGroup.GET("", co.GetPosts)
	postsGroup.GET("/:id", co.GetPost)
	postsGroup.POST("", co.CreatePost, m.Auth())
}
