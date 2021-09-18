package users

import (
	"go-rest-api/types/dto"
	"go-rest-api/types/entity"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	GetUsers(c echo.Context) error
	GetUser(c echo.Context) error
	CreateUser(c echo.Context) error
}

type UserService interface {
	GetAll(query dto.ListQuery) (dto.ListResult, error)
	GetById(id int) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
}

type UserRepository interface {
	FindAll(filter string) ([]*entity.User, error)
	FindById(id int) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
}
