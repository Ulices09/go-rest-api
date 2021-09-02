package posts

import (
	"github.com/labstack/echo/v4"
)

func NewPostRouter(e *echo.Echo, postController PostController) {
	postsGroup := e.Group("/posts")
	postsGroup.GET("", postController.GetPosts)
	postsGroup.GET("/:id", postController.GetPost)
	postsGroup.POST("", postController.CreatePost)
}
