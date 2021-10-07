package users_test

import (
	"go-rest-api/internal/core/dto"
	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/infrastructure/logger"
	"go-rest-api/internal/modules/users"
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

func (m *MockRepository) FindAll(filter string) ([]entity.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]entity.User), args.Error(1)
}

func (m *MockRepository) FindById(id int) (*entity.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*entity.User), args.Error(1)
}

func (m *MockRepository) Create(user users.CreateUserRequest) (*entity.User, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*entity.User), args.Error(1)
}

/*
  Test data
*/

var (
	createUserRequest = users.CreateUserRequest{
		Email:    "user1@email.com",
		Password: "my_hashed_password",
	}

	user1 = entity.User{
		Email:    "user1@email.com",
		Password: "my_hashed_password",
		Model: entity.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	user2 = entity.User{
		Email:    "user2@email.com",
		Password: "my_hashed_password",
		Model: entity.Model{
			ID:        2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	usersData = []entity.User{user1, user2}

	listQuery = dto.ListQuery{Filter: ""}
)

/*
  Test functions
*/
func TestGetAll(t *testing.T) {
	logger := logger.NewMockLogger(t)
	mockRepo := new(MockRepository)
	mockRepo.On("FindAll").Return(usersData, nil)

	service := users.NewUserService(mockRepo, logger)
	result, err := service.GetAll(listQuery)

	mockRepo.AssertExpectations(t)
	assert.Equal(t, 2, len(result.Data.([]entity.User)))
	assert.Nil(t, err)
}

func TestGetById(t *testing.T) {
	logger := logger.NewMockLogger(t)
	mockRepo := new(MockRepository)
	mockRepo.On("FindById").Return(&user1, nil)

	service := users.NewUserService(mockRepo, logger)
	user, err := service.GetById(1)

	mockRepo.AssertExpectations(t)
	assert.Equal(t, user1.ID, user.ID)
	assert.Equal(t, user1.Email, user.Email)
	assert.Equal(t, user1.Password, user.Password)
	assert.Nil(t, err)
}

func TestCreate(t *testing.T) {
	logger := logger.NewMockLogger(t)
	mockRepo := new(MockRepository)
	mockRepo.On("Create").Return(&user1, nil)

	service := users.NewUserService(mockRepo, logger)
	user, err := service.Create(createUserRequest)

	mockRepo.AssertExpectations(t)
	assert.Equal(t, user1.Email, user.Email)
	assert.Equal(t, user1.Password, "")
	assert.Nil(t, err)
}
