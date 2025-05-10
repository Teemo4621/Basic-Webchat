package entities

import (
	"time"

	"github.com/google/uuid"
)

type (
	Room struct {
		ID          uint      `gorm:"primaryKey autoIncrement" json:"id"`
		OwnerID     uint      `gorm:"not null" json:"owner_id"`
		Name        string    `gorm:"not null unique" json:"name"`
		Description string    `gorm:"not null" json:"description"`
		RoomCode    uuid.UUID `gorm:"not null unique" json:"room_code"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Messages    []Message `gorm:"foreignKey:RoomID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
	}

	RoomRepository interface {
		FindRoomsByUserId(id uint, page int, limit int, offset int) ([]Room, error)
		FindOneRoom(code string) (*Room, error)
		FindOneRoomByName(name string) (*Room, error)
		Create(room *Room) (*Room, error)
		Delete(id uint) error
		FindRoomsCountByUserId(id uint) (int64, error)
	}

	RoomUsecase interface {
		GetRoomsByUserId(id uint, page int, limit int) (RoomResponse, error)
		GetRoom(code string) (*Room, error)
		CreateRoom(room *Room) (*Room, error)
		JoinRoom(code string, userId uint) error
		LeaveRoom(code string, userID uint) error
		DeleteRoom(code string, userID uint) error
		GetRoomMembers(code string) ([]RoomMemberResponse, error)
	}

	RoomCreateRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	RoomResponse struct {
		Rooms     []Room `json:"rooms"`
		PageTotal int    `json:"page_total"`
	}

	RoomMemberResponse struct {
		RoomID     uint      `json:"room_id"`
		UserID     uint      `json:"user_id"`
		Username   string    `json:"username"`
		ProfileURL string    `json:"profile_url"`
		JoinedAt   time.Time `json:"joined_at"`
	}
)
