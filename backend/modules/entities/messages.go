package entities

import (
	"time"

	"github.com/google/uuid"
)

type (
	Message struct {
		ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
		MessageID uuid.UUID `gorm:"not null unique" json:"message_id"`
		RoomID    uint      `gorm:"not null" json:"room_id"`
		UserID    uint      `gorm:"not null" json:"user_id"`
		Content   string    `gorm:"type:text;not null" json:"content"`
		CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	}

	MessageRepository interface {
		Create(message *Message) (*Message, error)
		FindOneMessage(messageId uuid.UUID) (*Message, error)
		FindMessagesByRoomId(roomId uint, offset int, limit int) ([]Message, error)
		GetMessagesCount(roomId uint) (int64, error)
		Delete(messageId uuid.UUID) error
	}

	MessageUsecase interface {
		GetMessagesByRoomId(roomCode string, page int, limit int) (MessageResponse, error)
		SendMessage(roomCode string, userID uint, content string) (*Message, error)
		DeleteMessage(roomCode string, userID uint, messageID uuid.UUID) error
	}

	MessageResponse struct {
		Messages  []Message `json:"messages"`
		PageTotal int       `json:"page_total"`
	}

	MessageRequest struct {
		Content string `json:"content"`
	}
)
