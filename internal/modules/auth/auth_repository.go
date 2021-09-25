package auth

import (
	"context"
	"go-rest-api/ent"
	entUser "go-rest-api/ent/user"
	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/infrastructure/logger"
)

type repo struct {
	db     *ent.Client
	ctx    context.Context
	logger logger.Logger
}

func NewAuthRepository(db *ent.Client, logger logger.Logger) AuthRepository {
	ctx := context.Background()
	return &repo{db: db, ctx: ctx, logger: logger}
}

func (r *repo) GetUserByEmail(email string) (*entity.User, error) {
	result, err := r.db.User.Query().Where(entUser.Email(email)).Only(r.ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}

		r.logger.Errorw(err.Error(), "email", email)
		return nil, err
	}

	user := entity.NewUserFromSchema(result)
	return user, err
}
