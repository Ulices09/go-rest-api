package auth

import (
	"errors"
	"go-rest-api/internal/config"
	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/utils"
)

type service struct {
	authRepo AuthRepository
	config   config.Config
}

func NewAuthService(authRepo AuthRepository, config config.Config) AuthService {
	return &service{authRepo: authRepo, config: config}
}

func (s *service) Login(email, password string) (user *entity.User, token string, err error) {
	user, err = s.authRepo.GetUserByEmail(email)

	if err != nil {
		return
	}

	if user == nil {
		return nil, "", errors.New("incorrect credentials")
	}

	passwordOk := utils.CompareHash(password, user.Password)

	if !passwordOk {
		return nil, "", errors.New("incorrect credentials")
	}

	user.Password = ""

	token, err = utils.SignAuthJwt(*user, s.config.Jwt.Secret, s.config.Jwt.Expiration)

	if err != nil {
		return
	}

	return
}

func (s *service) Me(email string) (user *entity.User, err error) {
	user, err = s.authRepo.GetUserByEmail(email)

	if err != nil {
		return
	}

	user.Password = ""

	return
}
