package usecases

import (
	"errors"
	"log"

	"github.com/Teemo4621/Basic-Webchat/modules/entities"
	"github.com/google/uuid"
)

type roomUse struct {
	roomRepo       entities.RoomRepository
	roomMemberRepo entities.RoomMemberRepository
	messageRepo    entities.MessageRepository
	userRepo       entities.UserRepository
}

func NewRoomUsecase(roomRepo entities.RoomRepository, roomMemberRepo entities.RoomMemberRepository, messageRepo entities.MessageRepository, userRepo entities.UserRepository) entities.RoomUsecase {
	return &roomUse{
		roomRepo:       roomRepo,
		roomMemberRepo: roomMemberRepo,
		messageRepo:    messageRepo,
		userRepo:       userRepo,
	}
}

func (u *roomUse) GetRoomsByUserId(id uint, page int, limit int) (entities.RoomResponse, error) {
	totalCount, err := u.roomRepo.FindRoomsCountByUserId(id)
	if err != nil {
		log.Println(err.Error())
		return entities.RoomResponse{
			Rooms:     []entities.Room{},
			PageTotal: 0,
		}, errors.New("failed to get chats")
	}

	totalPages := int((totalCount + int64(limit) - 1) / int64(limit))

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	if page > totalPages {
		return entities.RoomResponse{
			Rooms:     []entities.Room{},
			PageTotal: totalPages,
		}, nil
	}

	offset := (page - 1) * limit

	if totalCount == 0 {
		return entities.RoomResponse{
			Rooms:     []entities.Room{},
			PageTotal: 0,
		}, nil
	}

	rooms, err := u.roomRepo.FindRoomsByUserId(id, page, limit, offset)
	if err != nil {
		log.Println(err.Error())
		return entities.RoomResponse{
			Rooms:     []entities.Room{},
			PageTotal: 0,
		}, errors.New("failed to get chats")
	}

	return entities.RoomResponse{
		Rooms:     rooms,
		PageTotal: totalPages,
	}, nil
}

func (u *roomUse) GetRoom(code string) (*entities.Room, error) {
	room, err := u.roomRepo.FindOneRoom(code)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (u *roomUse) CreateRoom(room *entities.Room) (*entities.Room, error) {
	checkRoomName, _ := u.roomRepo.FindOneRoomByName(room.Name)
	if checkRoomName != nil {
		return nil, errors.New("chat name already exists")
	}

	room.RoomCode = uuid.New()

	newRoom, err := u.roomRepo.Create(room)
	if err != nil {
		return nil, err
	}

	_, err = u.roomMemberRepo.Create(newRoom.ID, room.OwnerID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to create chat")
	}

	return newRoom, nil
}

func (u *roomUse) JoinRoom(code string, userID uint) error {
	room, err := u.GetRoom(code)
	if err != nil {
		return errors.New("room not found")
	}

	checkMember, _ := u.roomMemberRepo.FindMemberInRoom(room.ID, userID)
	if checkMember {
		return errors.New("you are already in this chat")
	}

	_, err = u.roomMemberRepo.Create(room.ID, userID)
	if err != nil {
		log.Println(err.Error())
		return errors.New("failed to join chat")
	}

	return nil
}

func (u *roomUse) LeaveRoom(code string, userID uint) error {
	room, err := u.GetRoom(code)
	if err != nil {
		return errors.New("room not found")
	}

	if room.OwnerID == userID {
		return errors.New("you are owner of this room")
	}

	if _, err := u.roomMemberRepo.FindMemberInRoom(room.ID, userID); err != nil {
		return errors.New("you are not in this room")
	}

	if err := u.roomMemberRepo.Delete(room.ID, userID); err != nil {
		log.Println(err.Error())
		return errors.New("failed to leave chat")
	}

	return nil
}

func (u *roomUse) DeleteRoom(code string, userID uint) error {
	room, err := u.GetRoom(code)
	if err != nil {
		return errors.New("room not found")
	}

	if _, err := u.roomMemberRepo.FindMemberInRoom(room.ID, userID); err != nil {
		return errors.New("you are not in this room")
	}

	if room.OwnerID != userID {
		return errors.New("you are not owner of this room")
	}

	if err := u.roomMemberRepo.Delete(room.ID, userID); err != nil {
		log.Println(err.Error())
		return errors.New("failed to delete room member")
	}

	if err := u.roomRepo.Delete(room.ID); err != nil {
		log.Println(err.Error())
		return errors.New("failed to delete room")
	}

	return nil
}

func (u *roomUse) GetRoomMembers(code string) ([]entities.RoomMemberResponse, error) {
	room, err := u.GetRoom(code)
	if err != nil {
		return nil, errors.New("room not found")
	}

	members, err := u.roomMemberRepo.FindByRoomId(room.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to get room members")
	}

	var roomMembers []entities.RoomMemberResponse

	for _, member := range members {
		user, err := u.userRepo.FindOneUserById(member.UserID)
		if err != nil {
			log.Println(err.Error())
			return nil, errors.New("failed to get user")
		}
		roomMembers = append(roomMembers, entities.RoomMemberResponse{
			RoomID:     room.ID,
			UserID:     member.UserID,
			Username:   user.Username,
			ProfileURL: user.ProfileURL,
			JoinedAt:   member.JoinedAt,
		})
	}

	return roomMembers, nil
}
