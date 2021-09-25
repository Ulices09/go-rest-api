package posts

import (
	"go-rest-api/internal/core/dto"
	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/core/errors"
	"go-rest-api/internal/infrastructure/logger"
)

type service struct {
	repo   PostRepository
	logger logger.Logger
}

func NewPostService(postRepo PostRepository, logger logger.Logger) PostService {
	return &service{repo: postRepo, logger: logger}
}

func (s *service) GetAll(query dto.PaginatedListQuery) (result dto.PaginationResult, err error) {
	posts, total, err := s.repo.FindAll(query.Filter, query.Skip, query.Take)

	if err != nil {
		s.logger.Errorw(err.Error(), "query", query)
		return
	}

	result = dto.NewPaginationResult(posts, total, query.Take)
	return
}

func (s *service) GetById(id int) (post *entity.Post, err error) {
	post, err = s.repo.FindById(id)

	if err != nil {
		s.logger.Errorw(err.Error(), "id", id)
		return
	}

	if post == nil {
		err = errors.NewNotFoundError("Post not found")
		return
	}

	return
}

func (s *service) Create(post *entity.Post, userId int) (newPost *entity.Post, err error) {
	newPost, err = s.repo.Create(post, userId)

	if err != nil {
		s.logger.Errorw(
			err.Error(),
			"post", post,
			"userId", userId,
		)
		return
	}

	return
}
