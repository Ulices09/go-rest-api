package entity

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtClaims struct {
	ID    int    `json:"id"`
	Email string `json:"emial"`
	jwt.StandardClaims
}

func NewCurrentUser(c echo.Context) JwtClaims {
	claims := c.Get("user").(*JwtClaims)
	return *claims
}
