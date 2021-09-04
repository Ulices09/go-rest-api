package posts

import (
	"go-rest-api/entity"

	"github.com/labstack/echo/v4"
)

type PostController interface {
	GetPosts(c echo.Context) error
	GetPost(c echo.Context) error
	CreatePost(c echo.Context) error
}

type PostService interface {
	GetAll() ([]*entity.Post, error)
	GetById(id int) (*entity.Post, error)
	Create(post *entity.Post) (*entity.Post, error)
}

type PostRepository interface {
	FindAll() ([]*entity.Post, error)
	FindById(id int) (*entity.Post, error)
	Create(post *entity.Post) (*entity.Post, error)
}
