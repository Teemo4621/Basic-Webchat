package usecases

import (
	"github.com/Teemo4621/Basic-Webchat/modules/entities"
)

type roomMemberUsecase struct {
	roomMemberRepo entities.RoomMemberRepository
}

func NewRoomMemberUsecase(roomMemberRepo entities.RoomMemberRepository) entities.RoomMemberUsecase {
	return &roomMemberUsecase{roomMemberRepo: roomMemberRepo}
}

func (u *roomMemberUsecase) FindMemberInRoom(roomID uint, userID uint) (bool, error) {
	return u.roomMemberRepo.FindMemberInRoom(roomID, userID)
}
