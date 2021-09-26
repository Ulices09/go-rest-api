package jwt_test

import (
	"testing"
	"time"

	"go-rest-api/internal/core/entity"
	"go-rest-api/internal/core/libs/jwt"

	"github.com/stretchr/testify/assert"
)

func TestSignAndCompare(t *testing.T) {
	secret := "my-secret"

	token, err := jwt.SignAuth(entity.User{}, secret, time.Now().Day())
	assert.NotEmpty(t, token)
	assert.Nil(t, err)

	result, err := jwt.Verify(token, secret)
	assert.NotNil(t, result)
	assert.Nil(t, err)

	result, err = jwt.Verify("random-token", secret)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}
