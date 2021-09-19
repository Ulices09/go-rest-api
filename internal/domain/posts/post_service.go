package posts

import (
	"go-rest-api/internal/core/dto"
	"go-rest-api/internal/core/entity"
)

type service struct {
	postRepo PostRepository
}

func NewPostService(postRepo PostRepository) PostService {
	return &service{postRepo}
}

func (s *service) GetAll(query dto.PaginatedListQuery) (result dto.PaginationResult, err error) {
	posts, total, err := s.postRepo.FindAll(query.Filter, query.Skip, query.Take)

	if err != nil {
		return
	}

	result = dto.NewPaginationResult(posts, total, query.Take)
	return
}

func (s *service) GetById(id int) (*entity.Post, error) {
	return s.postRepo.FindById(id)
}

func (s *service) Create(post *entity.Post, userId int) (*entity.Post, error) {
	return s.postRepo.Create(post, userId)
}
