package entities

import (
	"time"

	"github.com/Teemo4621/Basic-Webchat/configs"
)

type (
	AuthRepository interface {
		SaveRefreshToken(id uint, refreshToken string) error
		GetRefreshToken(id uint) (string, error)
	}

	AuthUsecase interface {
		Login(cfg *configs.Config, req *AuthLoginRequest) (*AuthLoginResponse, error)
		Register(req *AuthRegisterRequest) (*AuthRegisterResponse, error)
		Me(cfg *configs.Config, id uint) (*AuthMeResponse, error)
		RefreshToken(cfg *configs.Config, req *AuthRefreshTokenRequest) (*AuthRefreshTokenResponse, error)
	}

	AuthLoginResponse struct {
		Id           uint      `json:"id"`
		Username     string    `json:"username"`
		ProfileURL   string    `json:"profile_url"`
		CreatedAt    time.Time `json:"created_at"`
		AccessToken  string    `json:"access_token"`
		RefreshToken string    `json:"refresh_token"`
	}

	AuthRegisterResponse struct {
		Id        uint      `json:"id"`
		Username  string    `json:"username"`
		CreatedAt time.Time `json:"created_at"`
	}

	AuthMeResponse struct {
		Id         uint      `json:"id"`
		Username   string    `json:"username"`
		ProfileURL string    `json:"profile_url"`
		CreatedAt  time.Time `json:"created_at"`
	}

	AuthRefreshTokenResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	AuthLoginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	AuthRegisterRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	AuthRefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}
)
