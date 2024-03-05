package dto

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	jwt.StandardClaims
	Login string `json:"login"`
}
