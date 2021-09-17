package posts

import (
	"go-rest-api/types/entity"
)

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

func (s *service) Create(post *entity.Post, userId int) (*entity.Post, error) {
	return s.postRepo.Create(post, userId)
}
