package entities

import "github.com/golang-jwt/jwt/v4"

type (
	Jwtpassport struct {
		Id       uint   `json:"id"`
		Username string `json:"username"`
	}

	JwtClaim struct {
		Id       uint   `json:"id"`
		Username string `json:"username"`
		jwt.RegisteredClaims
	}
)
