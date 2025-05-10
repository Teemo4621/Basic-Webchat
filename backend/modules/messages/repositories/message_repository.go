package repositories

import (
	"github.com/Teemo4621/Basic-Webchat/modules/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type messageRepo struct {
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) entities.MessageRepository {
	return &messageRepo{DB: db}
}

func (m *messageRepo) Create(message *entities.Message) (*entities.Message, error) {
	if err := m.DB.Create(message).Error; err != nil {
		return nil, err
	}

	return message, nil
}

func (m *messageRepo) FindOneMessage(messageId uuid.UUID) (*entities.Message, error) {
	message := entities.Message{}
	if err := m.DB.First(&message, "message_id = ?", messageId).Error; err != nil {
		return nil, err
	}

	return &message, nil
}

func (m *messageRepo) GetMessagesCount(roomId uint) (int64, error) {
	var count int64
	if err := m.DB.Model(&entities.Message{}).Where("room_id = ?", roomId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (m *messageRepo) FindMessagesByRoomId(roomId uint, offset int, limit int) ([]entities.Message, error) {
	var messages []entities.Message
	if err := m.DB.Model(&entities.Message{}).
		Where("room_id = ?", roomId).
		Order("created_at desc").
		Offset(offset).
		Limit(limit).
		Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (m *messageRepo) Delete(messageId uuid.UUID) error {
	if err := m.DB.Delete(&entities.Message{}, "message_id = ?", messageId).Error; err != nil {
		return err
	}

	return nil
}
