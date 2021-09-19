package users

import (
	"context"
	"go-rest-api/ent"
	entUser "go-rest-api/ent/user"
	"go-rest-api/internal/core/entity"
)

type repo struct {
	db  *ent.Client
	ctx context.Context
}

func NewUserRepository(db *ent.Client) UserRepository {
	ctx := context.Background()
	return &repo{db: db, ctx: ctx}
}

func (r *repo) FindAll(filter string) ([]*entity.User, error) {
	results, err := r.db.User.
		Query().
		WithPosts().
		Select(
			entUser.FieldID,
			entUser.FieldEmail,
			entUser.FieldCreatedAt,
			entUser.FieldUpdatedAt,
		).
		Where(entUser.EmailContains(filter)).
		All(r.ctx)

	if err != nil {
		return nil, err
	}

	users := []*entity.User{}

	for _, result := range results {
		user := entity.NewUserFromSchema(result)
		users = append(users, user)
	}

	return users, err
}

func (r *repo) FindById(id int) (*entity.User, error) {
	result, err := r.db.User.
		Query().
		Select(
			entUser.FieldID,
			entUser.FieldEmail,
			entUser.FieldCreatedAt,
			entUser.FieldUpdatedAt,
		).
		WithPosts().
		Where(entUser.ID(id)).
		Only(r.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}

		return nil, err
	}

	user := entity.NewUserFromSchema(result)
	return user, err
}

func (r *repo) Create(user *entity.User) (*entity.User, error) {
	result, err := r.db.User.
		Create().
		SetEmail(user.Email).
		SetPassword(user.Password).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	newUser := entity.NewUserFromSchema(result)
	return newUser, err
}
