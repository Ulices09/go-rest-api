package auth

import (
	"context"
	"go-rest-api/ent"
	entUser "go-rest-api/ent/user"
	"go-rest-api/internal/types/entity"
)

type repo struct {
	db  *ent.Client
	ctx context.Context
}

func NewAuthRepository(db *ent.Client) AuthRepository {
	ctx := context.Background()
	return &repo{db: db, ctx: ctx}
}

func (r *repo) GetUserByEmail(email string) (*entity.User, error) {
	result, err := r.db.User.Query().Where(entUser.Email(email)).Only(r.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}

		return nil, err
	}

	user := entity.User{}
	user.MapFromSchema(result)

	return &user, err
}
