package posts

import (
	"context"
	"go-rest-api/ent"
	entPost "go-rest-api/ent/post"
	"go-rest-api/entity"
)

type repo struct {
	db  *ent.Client
	ctx context.Context
}

func NewPostRepository(db *ent.Client) PostRepository {
	ctx := context.Background()
	return &repo{db: db, ctx: ctx}
}

func (r *repo) FindAll() ([]*entity.Post, error) {
	results, err := r.db.Post.Query().WithUser().All(r.ctx)

	if err != nil {
		return nil, err
	}

	posts := []*entity.Post{}

	for _, result := range results {
		post := entity.Post{}
		post.MapFromSchema(result)

		posts = append(posts, &post)
	}

	return posts, err
}

func (r *repo) FindById(id int) (*entity.Post, error) {
	result, err := r.db.Post.Query().WithUser().Where(entPost.ID(id)).Only(r.ctx)

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

func (r *repo) Create(post *entity.Post) (*entity.Post, error) {
	// TODO: get userId from post
	result, err := r.db.Post.Create().SetTitle(post.Text).SetText(post.Text).SetUser(&ent.User{ID: 1}).Save(r.ctx)

	if err != nil {
		return nil, err
	}

	newPost := entity.Post{}
	newPost.MapFromSchema(result)

	return &newPost, err
}
