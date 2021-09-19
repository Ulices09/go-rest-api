package auth_test

import (
	"errors"
	"go-rest-api/internal/config"
	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/domain/auth"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
  Test objects
*/

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetUserByEmail(email string) (*entity.User, error) {
	args := m.Called()
	result := args.Get(0)

	if result == nil {
		return nil, args.Error(1)
	}

	return result.(*entity.User), args.Error(1)
}

/*
  Test data
*/

var (
	userSample = entity.User{
		Email:    "user1@email.com",
		Password: "$2a$10$KNX98LHqcr6uX58oA1vwAuwMPP2aOlFnc1ygBcMwBulRcroABJbDW",
		Model: entity.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	configSample = config.Config{
		Jwt: config.JwtConfig{
			Secret:     "mytestsecret",
			Expiration: 1000,
		},
	}

	errorSample = errors.New("sample error")
)

/*
  Test functions
*/
func TestLogin(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("GetUserByEmail").Return(&userSample, nil)

	email := "user1@email.com"
	password := "1234"

	service := auth.NewAuthService(mockRepo, configSample)
	user, token, err := service.Login(email, password)

	mockRepo.AssertExpectations(t)
	assert.Equal(t, email, user.Email)
	assert.NotEmpty(t, token)
	assert.Empty(t, user.Password)
	assert.Nil(t, err)
}

func TestLogin_IncorrectEmail(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("GetUserByEmail").Return(nil, nil)

	email := "user1@email.com"
	password := "1234"

	service := auth.NewAuthService(mockRepo, configSample)
	user, token, err := service.Login(email, password)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, user)
	assert.Empty(t, token)
	assert.NotNil(t, err)
}

func TestLogin_IncorrectPassword(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("GetUserByEmail").Return(&userSample, nil)

	email := "user1@email.com"
	password := "incorrect-password"

	service := auth.NewAuthService(mockRepo, configSample)
	user, token, err := service.Login(email, password)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, user)
	assert.Empty(t, token)
	assert.NotNil(t, err)
}

func TestMe(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("GetUserByEmail").Return(&userSample, nil)

	email := "user1@email.com"

	service := auth.NewAuthService(mockRepo, configSample)
	user, err := service.Me(email)

	mockRepo.AssertExpectations(t)
	assert.Equal(t, email, user.Email)
	assert.Empty(t, user.Password)
	assert.Nil(t, err)
}

func TestMe_Failure(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("GetUserByEmail").Return(nil, errorSample)

	email := "user1@email.com"

	service := auth.NewAuthService(mockRepo, configSample)
	user, err := service.Me(email)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}
