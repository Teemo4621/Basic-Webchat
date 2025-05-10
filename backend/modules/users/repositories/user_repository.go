package repositories

import (
	"github.com/Teemo4621/Basic-Webchat/modules/entities"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) entities.UserRepository {
	return &UserRepo{DB: db}
}

func (r *UserRepo) FindOneUser(username string) (*entities.User, error) {
	user := entities.User{}

	if err := r.DB.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) FindOneUserById(id uint) (*entities.User, error) {
	user := entities.User{}

	if err := r.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) FindAllUsers() ([]entities.User, error) {
	users := []entities.User{}

	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepo) Create(user *entities.User) (*entities.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) Update(user *entities.User) (*entities.User, error) {
	if err := r.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
