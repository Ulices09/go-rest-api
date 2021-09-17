package posts

import (
	"context"
	"go-rest-api/ent"
	entPost "go-rest-api/ent/post"
	"go-rest-api/types/entity"
)

type repo struct {
	db  *ent.Client
	ctx context.Context
}

func NewPostRepository(db *ent.Client) PostRepository {
	ctx := context.Background()
	return &repo{db: db, ctx: ctx}
}

func (r *repo) FindAll(skip, take int) ([]*entity.Post, int, error) {
	query := r.db.Post.
		Query().
		WithUser().
		Clone()

	count, err := query.Count(r.ctx)

	if err != nil {
		return nil, 0, err
	}

	results, err := query.
		Offset(skip).
		Limit(take).
		All(r.ctx)

	if err != nil {
		return nil, 0, err
	}

	posts := []*entity.Post{}

	for _, result := range results {
		post := entity.Post{}
		post.MapFromSchema(result)

		posts = append(posts, &post)
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

		return nil, err
	}

	post := entity.Post{}
	post.MapFromSchema(result)

	return &post, err
}

func (r *repo) Create(post *entity.Post, userId int) (*entity.Post, error) {
	result, err := r.db.Post.
		Create().
		SetTitle(post.Text).
		SetText(post.Text).
		SetUser(&ent.User{ID: userId}).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	newPost := entity.Post{}
	newPost.MapFromSchema(result)

	return &newPost, err
}
