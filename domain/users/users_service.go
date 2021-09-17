package users

import (
	"go-rest-api/types/entity"
	"go-rest-api/utils"
)

type service struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &service{userRepo}
}

func (s *service) GetAll() ([]*entity.User, error) {
	return s.userRepo.FindAll()
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
