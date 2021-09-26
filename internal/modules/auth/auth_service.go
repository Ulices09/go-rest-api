package auth

import (
	"go-rest-api/internal/config"
	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/core/errors"
	"go-rest-api/internal/core/libs/hash"
	"go-rest-api/internal/core/libs/jwt"
	"go-rest-api/internal/infrastructure/logger"
)

type service struct {
	repo   AuthRepository
	config config.Config
	logger logger.Logger
}

func NewAuthService(
	authRepo AuthRepository,
	config config.Config,
	logger logger.Logger) AuthService {
	return &service{repo: authRepo, config: config, logger: logger}
}

func (s *service) Login(email, password string) (user *entity.User, token string, err error) {
	user, err = s.repo.GetUserByEmail(email)

	if err != nil {
		s.logger.Errorw(err.Error(), "email", email)
		return
	}

	if user == nil {
		err = errors.NewBadRequestError("Incorrect credentials")
		return
	}

	passwordOk := hash.Compare(password, user.Password)

	if !passwordOk {
		err = errors.NewBadRequestError("Incorrect credentials")
		return
	}

	user.Password = ""

	token, err = jwt.SignAuth(*user, s.config.Jwt.Secret, s.config.Jwt.Expiration)

	if err != nil {
		s.logger.Errorw(err.Error(), "email", email)
		return
	}

	return
}

func (s *service) Me(email string) (user *entity.User, err error) {
	user, err = s.repo.GetUserByEmail(email)

	if err != nil {
		s.logger.Errorw(err.Error(), "email", email)
		return
	}

	user.Password = ""

	return
}
