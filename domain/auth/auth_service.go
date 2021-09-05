package auth

import (
	"errors"
	"go-rest-api/entity"
	"go-rest-api/utils"
)

type service struct {
	authRepo AuthRepository
}

func NewAuthService(authRepo AuthRepository) AuthService {
	return &service{authRepo}
}

func (s *service) Login(email, password string) (user *entity.User, err error) {
	user, err = s.authRepo.GetUser(email)

	if err != nil {
		return
	}

	if user == nil {
		return nil, errors.New("incorrect credentials")
	}

	passwordOk := utils.CompareHash(password, user.Password)

	if !passwordOk {
		return nil, errors.New("incorrect credentials")
	}

	user.Password = ""

	return
}
