package posts

import (
	"errors"
	"go-rest-api/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	FindAll() ([]entity.Post, error)
	FindById(id int) (*entity.Post, error)
	Create(post *entity.Post) (*entity.Post, error)
}

type repo struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &repo{db}
}

func (r *repo) FindAll() ([]entity.Post, error) {
	var posts []entity.Post

	if result := r.db.Find(&posts); result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

func (r *repo) FindById(id int) (*entity.Post, error) {
	var post entity.Post

	if result := r.db.First(&post, id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &post, nil
}

func (r *repo) Create(post *entity.Post) (*entity.Post, error) {
	if result := r.db.Create(&post); result.Error != nil {
		return nil, result.Error
	}

	return post, nil
}
