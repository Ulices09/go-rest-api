package hash_test

import (
	"testing"

	"go-rest-api/internal/core/libs/hash"

	"github.com/stretchr/testify/assert"
)

func TestHashAndCompare(t *testing.T) {
	value := "sample-value"

	hashedValue, err := hash.Hash(value)
	assert.NotEmpty(t, hashedValue)
	assert.Nil(t, err)

	ok := hash.Compare(value, hashedValue)
	assert.True(t, ok)

	ok = hash.Compare("random-value", hashedValue)
	assert.False(t, ok)
}
