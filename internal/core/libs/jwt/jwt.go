package jwt

import (
	"fmt"
	"go-rest-api/internal/core/entity"
	"time"

	_jwt "github.com/golang-jwt/jwt"
)

func SignAuth(user entity.User, secret string, expiration int) (string, error) {
	expiredTime := time.Now().Add(time.Minute * time.Duration(expiration))

	claims := &entity.JwtClaims{
		ID:     user.ID,
		Email:  user.Email,
		RoleId: user.Role.ID,
		StandardClaims: _jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
		},
	}

	token := _jwt.NewWithClaims(_jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))

	return signedToken, err
}

func Verify(value string, secret string) (*_jwt.Token, error) {
	token, err := _jwt.ParseWithClaims(
		value,
		&entity.JwtClaims{},
		func(token *_jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*_jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		},
	)

	return token, err
}
