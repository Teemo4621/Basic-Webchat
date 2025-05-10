package repositories

import (
	"github.com/Teemo4621/Basic-Webchat/modules/entities"
	"gorm.io/gorm"
)

type roomRepo struct {
	DB *gorm.DB
}

func NewRoomRepository(db *gorm.DB) entities.RoomRepository {
	return &roomRepo{DB: db}
}

func (r *roomRepo) FindOneRoom(code string) (*entities.Room, error) {
	room := entities.Room{}

	if err := r.DB.First(&room, "room_code = ?", code).Error; err != nil {
		return nil, err
	}

	return &room, nil
}

func (r *roomRepo) FindOneRoomByName(name string) (*entities.Room, error) {
	room := entities.Room{}

	if err := r.DB.First(&room, "name = ?", name).Error; err != nil {
		return nil, err
	}

	return &room, nil
}

func (r *roomRepo) FindRoomsCountByUserId(id uint) (int64, error) {
	var count int64
	if err := r.DB.Model(&entities.RoomMember{}).Where("user_id = ?", id).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *roomRepo) FindRoomsByUserId(id uint, page int, limit int, offset int) ([]entities.Room, error) {
	var rooms []entities.Room
	var roomMembers []entities.RoomMember

	if err := r.DB.Model(&entities.RoomMember{}).Where("user_id = ?", id).Offset(offset).Limit(limit).Find(&roomMembers).Error; err != nil {
		return nil, err
	}

	for _, roomMember := range roomMembers {
		room := entities.Room{}
		if err := r.DB.First(&room, "id = ?", roomMember.RoomID).Error; err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r *roomRepo) Create(room *entities.Room) (*entities.Room, error) {
	if err := r.DB.Create(room).Error; err != nil {
		return nil, err
	}

	return room, nil
}

func (r *roomRepo) Delete(id uint) error {
	if err := r.DB.Delete(&entities.Room{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
