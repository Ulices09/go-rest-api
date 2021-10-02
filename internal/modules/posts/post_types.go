package posts

import (
	"go-rest-api/internal/core/dto"
	"go-rest-api/internal/core/entity"

	"github.com/labstack/echo/v4"
)

/*
  Interfaces
*/

type PostController interface {
	GetPosts(c echo.Context) error
	GetPost(c echo.Context) error
	CreatePost(c echo.Context) error
}

type PostService interface {
	GetAll(query dto.PaginatedListQuery) (dto.PaginationResult, error)
	GetById(id int) (*entity.Post, error)
	Create(post CreatePostRequest, userId int) (*entity.Post, error)
}

type PostRepository interface {
	FindAll(filter string, skip int, take int) ([]*entity.Post, int, error)
	FindById(id int) (*entity.Post, error)
	Create(post CreatePostRequest, userId int) (*entity.Post, error)
}

/*
  DTOs
*/

type CreatePostRequest struct {
	Title string `json:"title" validate:"required"`
	Text  string `json:"text" validate:"required"`
}
