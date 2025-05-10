package repositories

import (
	"github.com/Teemo4621/Basic-Webchat/modules/entities"
	"gorm.io/gorm"
)

type AuthRepo struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) entities.AuthRepository {
	return &AuthRepo{DB: db}
}

func (r *AuthRepo) SaveRefreshToken(id uint, refreshToken string) error {
	if err := r.DB.Model(&entities.User{}).Where("id = ?", id).Update("refresh_token", refreshToken).Error; err != nil {
		return err
	}
	return nil
}

func (r *AuthRepo) GetRefreshToken(id uint) (string, error) {
	var user entities.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return "", err
	}
	return user.RefreshToken, nil
}
