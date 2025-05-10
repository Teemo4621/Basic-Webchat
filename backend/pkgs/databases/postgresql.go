package databases

import (
	"log"

	"github.com/Teemo4621/Basic-Webchat/configs"
	"github.com/Teemo4621/Basic-Webchat/modules/entities"
	"github.com/Teemo4621/Basic-Webchat/pkgs/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection(config configs.Config) (*gorm.DB, error) {
	url, err := utils.ConnectionURLBuilder("postgres", config)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("PostgreSQL database has been connected 📦")
	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&entities.User{}, &entities.Room{}, &entities.RoomMember{}, &entities.Message{})
}
