package posts

import (
	"go-rest-api/types/dto"
	"go-rest-api/types/entity"

	"github.com/labstack/echo/v4"
)

type PostController interface {
	GetPosts(c echo.Context) error
	GetPost(c echo.Context) error
	CreatePost(c echo.Context) error
}

type PostService interface {
	GetAll(query dto.PaginationQuery) (dto.PaginationResult, error)
	GetById(id int) (*entity.Post, error)
	Create(post *entity.Post, userId int) (*entity.Post, error)
}

type PostRepository interface {
	FindAll(skip, take int) ([]*entity.Post, int, error)
	FindById(id int) (*entity.Post, error)
	Create(post *entity.Post, userId int) (*entity.Post, error)
}
