package entities

import "time"

type (
	RoomMember struct {
		RoomID   uint      `gorm:"primaryKey;not null" json:"room_id"`
		UserID   uint      `gorm:"primaryKey;not null" json:"user_id"`
		JoinedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"joined_at"`
	}

	RoomMemberRepository interface {
		FindByRoomId(id uint) ([]RoomMember, error)
		FindMemberInRoom(roomID uint, userID uint) (bool, error)
		Create(id uint, userId uint) (*RoomMember, error)
		Delete(id uint, userID uint) error
	}

	RoomMemberUsecase interface {
		FindMemberInRoom(roomID uint, userID uint) (bool, error)
	}
)
