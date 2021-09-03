package posts

import (
	"go-rest-api/entity"
)

type PostService interface {
	GetAll() ([]*entity.Post, error)
	GetById(id int) (*entity.Post, error)
	Create(post *entity.Post) (*entity.Post, error)
}

type service struct {
	postRepo PostRepository
}

func NewPostService(postRepo PostRepository) PostService {
	return &service{postRepo}
}

func (s *service) GetAll() ([]*entity.Post, error) {
	return s.postRepo.FindAll()
}

func (s *service) GetById(id int) (*entity.Post, error) {
	return s.postRepo.FindById(id)
}

func (s *service) Create(post *entity.Post) (*entity.Post, error) {
	return s.postRepo.Create(post)
}
