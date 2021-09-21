package auth

import (
	"go-rest-api/internal/config"
	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/core/errors"
	"go-rest-api/internal/core/utils"
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
		err = errors.NewBadRequestError("Incorrect credentials")
		return
	}

	passwordOk := utils.CompareHash(password, user.Password)

	if !passwordOk {
		err = errors.NewBadRequestError("Incorrect credentials")
		return
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
