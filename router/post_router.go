package router

import (
	"go-rest-api/controller"

	"github.com/labstack/echo/v4"
)

func NewPostRouter(e *echo.Echo, postController controller.PostController) {
	postsGroup := e.Group("/posts")
	postsGroup.GET("", postController.GetPosts)
	postsGroup.GET("/:id", postController.GetPost)
	postsGroup.POST("", postController.CreatePost)
}
