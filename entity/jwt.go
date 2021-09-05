package entity

import "github.com/golang-jwt/jwt"

type Claims struct {
	ID    int    `json:"id"`
	Email string `json:"emial"`
	jwt.StandardClaims
}
