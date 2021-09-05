package auth

import (
	"go-rest-api/entity"

	"github.com/labstack/echo/v4"
)

/*
  Interfaces
*/

type AuthController interface {
	Login(c echo.Context) error
}

type AuthService interface {
	Login(email, password string) (*entity.User, error)
}

type AuthRepository interface {
	GetUser(email string) (*entity.User, error)
}

/*
  DTOs
*/

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}