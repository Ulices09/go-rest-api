package users

import (
	"go-rest-api/internal/core/dto"
	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/utils"
)

type service struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &service{userRepo}
}

func (s *service) GetAll(query dto.ListQuery) (result dto.ListResult, err error) {
	users, err := s.userRepo.FindAll(query.Filter)

	if err != nil {
		return
	}

	result = dto.ListResult{
		Data: users,
	}

	return
}

func (s *service) GetById(id int) (*entity.User, error) {
	return s.userRepo.FindById(id)
}

func (s *service) Create(user *entity.User) (newUser *entity.User, err error) {
	hashedPassword, err := utils.Hash(user.Password)

	if err != nil {
		return
	}

	user.Password = hashedPassword
	newUser, err = s.userRepo.Create(user)

	if err != nil {
		return
	}

	newUser.Password = ""

	return
}
