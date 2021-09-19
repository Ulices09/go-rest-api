package posts_test

import (
	"go-rest-api/internal/core/dto"
	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/domain/posts"
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

func (m *MockRepository) FindAll(filter string, skip int, take int) ([]*entity.Post, int, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]*entity.Post), args.Int(1), args.Error(2)
}

func (m *MockRepository) FindById(id int) (*entity.Post, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (m *MockRepository) Create(post *entity.Post, userId int) (*entity.Post, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

/*
  Test data
*/

var (
	post1 = entity.Post{
		Title: "My first post",
		Text:  "This is my first post",
		Model: entity.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	post2 = entity.Post{
		Title: "My second post",
		Text:  "This is my second post",
		Model: entity.Model{
			ID:        2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	postsData = []*entity.Post{&post1, &post2}

	paginatedListQuery = dto.PaginatedListQuery{
		ListQuery: dto.ListQuery{
			Filter: "",
		},
		PaginationQuery: dto.PaginationQuery{
			Page: 1,
			Skip: 0,
			Take: 10,
		},
	}
)

/*
  Test functions
*/
func TestGetAll(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("FindAll").Return(postsData, 100, nil)

	service := posts.NewPostService(mockRepo)
	result, err := service.GetAll(paginatedListQuery)

	mockRepo.AssertExpectations(t)
	assert.Equal(t, 2, len(result.Data.([]*entity.Post)))
	assert.Nil(t, err)
}

func TestGetById(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("FindById").Return(&post1, nil)

	service := posts.NewPostService(mockRepo)
	post, err := service.GetById(1)

	mockRepo.AssertExpectations(t)
	assert.Equal(t, post1.ID, post.ID)
	assert.Equal(t, post1.Title, post.Title)
	assert.Equal(t, post1.Text, post.Text)
	assert.Nil(t, err)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("Create").Return(&post1, nil)

	service := posts.NewPostService(mockRepo)
	post, err := service.Create(&post1, 1)

	mockRepo.AssertExpectations(t)
	assert.Equal(t, post1.Title, post.Title)
	assert.Equal(t, post1.Text, post.Text)
	assert.Nil(t, err)
}
