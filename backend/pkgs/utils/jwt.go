package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/Teemo4621/Basic-Webchat/configs"
	"github.com/Teemo4621/Basic-Webchat/modules/entities"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GenerateAccessToken(cfg *configs.Config, req *entities.Jwtpassport) (string, error) {
	claims := entities.JwtClaim{
		Id:       req.Id,
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(cfg.JWT.Expire))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "access_token",
			Subject:   "user_access_token",
			ID:        uuid.NewString(),
			Audience:  []string{"user"},
		},
	}

	mySignKey := cfg.JWT.Secret

	if mySignKey == "" {
		return "", errors.New("error, jwt secret is empty")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(mySignKey))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(cfg *configs.Config, req *entities.Jwtpassport) (string, error) {
	claims := entities.JwtClaim{
		Id:       req.Id,
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(cfg.JWT.RefreshExpire))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "refresh_token",
			Subject:   "user_refresh_token",
			ID:        uuid.NewString(),
			Audience:  []string{"user"},
		},
	}

	mySignKey := cfg.JWT.RefreshSecret

	if mySignKey == "" {
		return "", errors.New("error, jwt refresh secret is empty")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(mySignKey))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenString, nil
}

func ValidateAccessToken(cfg *configs.Config, tokenString string) (*jwt.Token, error) {
	mySignKey := cfg.JWT.Secret

	if mySignKey == "" {
		return nil, errors.New("error, jwt secret is empty")
	}

	token, err := jwt.ParseWithClaims(tokenString, &entities.JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySignKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func ParseAccessToken(cfg *configs.Config, tokenString string) (*entities.JwtClaim, error) {
	token, err := ValidateAccessToken(cfg, tokenString)
	if err != nil {
		return nil, err
	}

	claim, ok := token.Claims.(*entities.JwtClaim)
	if !ok {
		return nil, errors.New("error, jwt claim is not valid")
	}

	return claim, nil
}

func ValidateRefreshToken(cfg *configs.Config, tokenString string) (*jwt.Token, error) {
	mySignKey := cfg.JWT.RefreshSecret

	if mySignKey == "" {
		return nil, errors.New("error, jwt refresh secret is empty")
	}

	token, err := jwt.ParseWithClaims(tokenString, &entities.JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySignKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func ParseRefreshToken(cfg *configs.Config, tokenString string) (*entities.JwtClaim, error) {
	token, err := ValidateRefreshToken(cfg, tokenString)
	if err != nil {
		return nil, err
	}

	claim, ok := token.Claims.(*entities.JwtClaim)
	if !ok {
		return nil, errors.New("error, jwt claim is not valid")
	}

	return claim, nil
}
