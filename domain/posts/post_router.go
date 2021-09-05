package posts

import (
	"github.com/labstack/echo/v4"
)

func NewPostRouter(e *echo.Echo, co PostController) {
	postsGroup := e.Group("/posts")
	postsGroup.GET("", co.GetPosts)
	postsGroup.GET("/:id", co.GetPost)
	postsGroup.POST("", co.CreatePost)
}
