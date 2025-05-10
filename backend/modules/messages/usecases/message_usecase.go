package usecases

import (
	"errors"
	"log"

	"github.com/Teemo4621/Basic-Webchat/modules/entities"
	"github.com/google/uuid"
)

type messageUse struct {
	messageRepo    entities.MessageRepository
	roomRepo       entities.RoomRepository
	roomMemberRepo entities.RoomMemberRepository
}

func NewMessageUsecase(messageRepo entities.MessageRepository, roomRepo entities.RoomRepository, roomMemberRepo entities.RoomMemberRepository) entities.MessageUsecase {
	return &messageUse{messageRepo: messageRepo, roomRepo: roomRepo, roomMemberRepo: roomMemberRepo}
}

func (u *messageUse) GetMessagesByRoomId(roomCode string, page int, limit int) (entities.MessageResponse, error) {
	room, err := u.roomRepo.FindOneRoom(roomCode)
	if err != nil {
		return entities.MessageResponse{}, errors.New("room not found")
	}

	totalCount, err := u.messageRepo.GetMessagesCount(room.ID)
	if err != nil {
		return entities.MessageResponse{}, errors.New("failed to get messages count")
	}

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	totalPages := int((totalCount + int64(limit) - 1) / int64(limit))

	if page > totalPages {
		return entities.MessageResponse{
			Messages:  []entities.Message{},
			PageTotal: totalPages,
		}, nil
	}

	offset := (page - 1) * limit

	messages, err := u.messageRepo.FindMessagesByRoomId(room.ID, int(offset), limit)
	if err != nil {
		return entities.MessageResponse{}, errors.New("failed to get messages")
	}

	return entities.MessageResponse{
		Messages:  messages,
		PageTotal: totalPages,
	}, nil
}

func (u *messageUse) SendMessage(roomCode string, userID uint, content string) (*entities.Message, error) {
	room, err := u.roomRepo.FindOneRoom(roomCode)
	if err != nil {
		return nil, errors.New("room not found")
	}

	if _, err := u.roomMemberRepo.FindMemberInRoom(room.ID, userID); err != nil {
		return nil, errors.New("you are not in this room")
	}

	message := entities.Message{
		MessageID: uuid.New(),
		RoomID:    room.ID,
		UserID:    userID,
		Content:   content,
	}

	newMessage, err := u.messageRepo.Create(&message)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("failed to send message")
	}

	return newMessage, nil
}

func (u *messageUse) DeleteMessage(roomCode string, userID uint, messageID uuid.UUID) error {
	room, err := u.roomRepo.FindOneRoom(roomCode)
	if err != nil {
		return errors.New("room not found")
	}

	message, err := u.messageRepo.FindOneMessage(messageID)
	if err != nil {
		return errors.New("message not found")
	}

	if message.RoomID != room.ID {
		return errors.New("message not found")
	}

	if message.UserID != userID {
		return errors.New("you are not owner of this message")
	}

	if _, err := u.roomMemberRepo.FindMemberInRoom(room.ID, userID); err != nil {
		return errors.New("you are not in this room")
	}

	if err := u.messageRepo.Delete(messageID); err != nil {
		log.Println(err.Error())
		return errors.New("failed to delete message")
	}

	return nil
}
