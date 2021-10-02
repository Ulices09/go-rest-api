package posts

import (
	"context"
	"go-rest-api/ent"
	entPost "go-rest-api/ent/post"
	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/infrastructure/logger"
)

type repo struct {
	db     *ent.Client
	ctx    context.Context
	logger logger.Logger
}

func NewPostRepository(db *ent.Client, logger logger.Logger) PostRepository {
	ctx := context.Background()
	return &repo{db: db, ctx: ctx, logger: logger}
}

func (r *repo) FindAll(filter string, skip int, take int) ([]*entity.Post, int, error) {
	query := r.db.Post.
		Query().
		WithUser().
		Where(entPost.TitleContains(filter)).
		Clone()

	count, err := query.Count(r.ctx)

	if err != nil {
		r.logger.Errorw(
			err.Error(),
			"filter", filter,
			"skip", skip,
			"take", take,
		)
		return nil, 0, err
	}

	results, err := query.
		Offset(skip).
		Limit(take).
		All(r.ctx)

	if err != nil {
		r.logger.Errorw(
			err.Error(),
			"filter", filter,
			"skip", skip,
			"take", take,
		)
		return nil, 0, err
	}

	posts := []*entity.Post{}

	for _, result := range results {
		post := entity.NewPostFromSchema(result)
		posts = append(posts, post)
	}

	return posts, count, err
}

func (r *repo) FindById(id int) (*entity.Post, error) {
	result, err := r.db.Post.
		Query().
		WithUser().
		Where(entPost.ID(id)).
		Only(r.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}

		r.logger.Errorw(err.Error(), "id", id)
		return nil, err
	}

	post := entity.NewPostFromSchema(result)
	return post, err
}

func (r *repo) Create(post CreatePostRequest, userId int) (*entity.Post, error) {
	result, err := r.db.Post.
		Create().
		SetTitle(post.Title).
		SetText(post.Text).
		SetUser(&ent.User{ID: userId}).
		Save(r.ctx)

	if err != nil {
		r.logger.Errorw(
			err.Error(),
			"post", post,
			"userId", userId,
		)
		return nil, err
	}

	newPost := entity.NewPostFromSchema(result)
	return newPost, err
}
