package user

import "github.com/golang-jwt/jwt/v5"

type TokenClaims struct {
	jwt.RegisteredClaims
	UserId string `json:"userId"`
	Role   string `json:"role"`
}
