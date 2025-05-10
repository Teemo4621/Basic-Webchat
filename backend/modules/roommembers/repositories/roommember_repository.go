package repositories

import (
	"github.com/Teemo4621/Basic-Webchat/modules/entities"
	"gorm.io/gorm"
)

type roomMemberRepo struct {
	DB *gorm.DB
}

func NewRoomMemberRepository(db *gorm.DB) entities.RoomMemberRepository {
	return &roomMemberRepo{DB: db}
}

func (r *roomMemberRepo) Create(roomID uint, userID uint) (*entities.RoomMember, error) {
	roomMember := entities.RoomMember{
		RoomID: roomID,
		UserID: userID,
	}

	if err := r.DB.Create(&roomMember).Error; err != nil {
		return nil, err
	}

	return &roomMember, nil
}

func (r *roomMemberRepo) FindMemberInRoom(roomID uint, userID uint) (bool, error) {
	roomMember := entities.RoomMember{}

	if err := r.DB.First(&roomMember, "room_id = ? AND user_id = ?", roomID, userID).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *roomMemberRepo) Delete(roomID uint, userID uint) error {
	if err := r.DB.Delete(&entities.RoomMember{}, "room_id = ? AND user_id = ?", roomID, userID).Error; err != nil {
		return err
	}

	return nil
}

func (r *roomMemberRepo) FindByRoomId(id uint) ([]entities.RoomMember, error) {
	roomMembers := []entities.RoomMember{}

	if err := r.DB.Where("room_id = ?", id).Find(&roomMembers).Error; err != nil {
		return nil, err
	}

	return roomMembers, nil
}

func (r *roomMemberRepo) FindByUserId(id uint) ([]entities.RoomMember, error) {
	roomMembers := []entities.RoomMember{}

	if err := r.DB.Where("user_id = ?", id).Find(&roomMembers).Error; err != nil {
		return nil, err
	}

	return roomMembers, nil
}
