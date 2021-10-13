package jwt_test

import (
	"testing"
	"time"

	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/core/libs/jwt"

	"github.com/stretchr/testify/assert"
)

var (
	secret = "my-secret"

	user = entity.User{
		Model: entity.Model{
			ID: 1,
		},
		Email: "email@email.com",
		Role: &entity.Role{
			Model: entity.Model{
				ID: 1,
			},
		},
	}
)

func TestSignAndVerify(t *testing.T) {

	token, err := jwt.SignAuth(user, secret, time.Now().Day())
	assert.NotEmpty(t, token)
	assert.Nil(t, err)

	result, err := jwt.Verify(token, secret)
	assert.NotNil(t, result)
	assert.Nil(t, err)

	result, err = jwt.Verify("random-token", secret)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}
