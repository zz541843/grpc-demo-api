package models

import (
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	ID       uint
	Username string
	Password string
	jwt.StandardClaims
}
