package utils

import (
	"fmt"
	"go-rest-api/internal/core/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

func SignAuthJwt(user entity.User, secret string, expiration int) (string, error) {
	expiredTime := time.Now().Add(time.Minute * time.Duration(expiration))

	claims := &entity.Claims{
		ID:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))

	return signedToken, err
}

func VerifyJwt(value string, secret string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(
		value,
		&entity.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		},
	)

	return token, err
}
