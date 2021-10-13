package users

import (
	"go-rest-api/internal/core/dto"
	"go-rest-api/internal/core/entity"

	"github.com/labstack/echo/v4"
)

/*
  Interfaces
*/

type UserController interface {
	GetUsers(c echo.Context) error
	GetUser(c echo.Context) error
	CreateUser(c echo.Context) error
}

type UserService interface {
	GetAll(query dto.ListQuery) (dto.ListResult, error)
	GetById(id int) (*entity.User, error)
	Create(user CreateUserRequest) (*entity.User, error)
}

type UserRepository interface {
	FindAll(filter string) ([]entity.User, error)
	FindById(id int) (*entity.User, error)
	Create(user CreateUserRequest) (*entity.User, error)
}

/*
  DTOs
*/

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	RoleId   int    `json:"roleId" validate:"required"`
}
