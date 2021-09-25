package users

import (
	"go-rest-api/internal/core/dto"
	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/core/errors"
	"go-rest-api/internal/core/utils"
	"go-rest-api/internal/infrastructure/logger"
)

type service struct {
	repo   UserRepository
	logger logger.Logger
}

func NewUserService(userRepo UserRepository, logger logger.Logger) UserService {
	return &service{repo: userRepo, logger: logger}
}

func (s *service) GetAll(query dto.ListQuery) (result dto.ListResult, err error) {
	users, err := s.repo.FindAll(query.Filter)

	if err != nil {
		s.logger.Errorw(err.Error(), "query", query)
		return
	}

	result = dto.ListResult{
		Data: users,
	}

	return
}

func (s *service) GetById(id int) (user *entity.User, err error) {
	user, err = s.repo.FindById(id)

	if err != nil {
		s.logger.Errorw(err.Error(), "id", id)
		return
	}

	if user == nil {
		err = errors.NewNotFoundError("Post not found")
		return
	}

	return
}

func (s *service) Create(user *entity.User) (newUser *entity.User, err error) {
	hashedPassword, err := utils.Hash(user.Password)

	if err != nil {
		userToLog := entity.NewUserToLog(*user)
		s.logger.Errorw(err.Error(), "user", userToLog)
		return
	}

	user.Password = hashedPassword
	newUser, err = s.repo.Create(user)

	if err != nil {
		userToLog := entity.NewUserToLog(*user)
		s.logger.Errorw(err.Error(), "user", userToLog)
		return
	}

	newUser.Password = ""

	return
}
