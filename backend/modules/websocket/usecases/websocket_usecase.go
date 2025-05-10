package usecases

import (
	"fmt"
	"log"
	"sync"

	"github.com/Teemo4621/Basic-Webchat/modules/entities"
	"github.com/gofiber/contrib/websocket"
)

type websocketUse struct {
	UserRepo       entities.UserRepository
	RoomRepo       entities.RoomRepository
	RoomMemberRepo entities.RoomMemberRepository
	mu             sync.RWMutex
	connections    map[string]map[string]*websocket.Conn
}

func NewWebSocketUsecase(userRepo entities.UserRepository, roomRepo entities.RoomRepository, roomMemberRepo entities.RoomMemberRepository) entities.WebSocketUsecase {
	return &websocketUse{
		UserRepo:       userRepo,
		RoomRepo:       roomRepo,
		RoomMemberRepo: roomMemberRepo,
		connections:    make(map[string]map[string]*websocket.Conn),
	}
}

func (u *websocketUse) AddConnection(roomCode string, userID uint, conn *websocket.Conn) error {
	u.mu.Lock()

	log.Printf("AddConnection: userID=%d, roomCode=%s", userID, roomCode)

	room, err := u.RoomRepo.FindOneRoom(roomCode)
	if err != nil {
		log.Printf("FindOneRoom error: %v", err)
		return err
	}

	if _, err := u.RoomMemberRepo.FindMemberInRoom(room.ID, userID); err != nil {
		log.Printf("User %d is not a member of room %s", userID, roomCode)
		return err
	}

	user, err := u.UserRepo.FindOneUserById(userID)
	if err != nil {
		log.Printf("FindOneUserById error: %v", err)
		return err
	}

	defer func() {
		u.mu.Unlock()
		//users := u.GetUsersInRoom(roomCode)
		// for _, username := range users {
		// 	u.BroadcastMessage(roomCode, entities.WebSocketRequest{
		// 		Method: "system",
		// 		Data: map[string]interface{}{
		// 			"message":  fmt.Sprintf("%s ได้เข้าร่วมห้อง", username),
		// 			"username": username,
		// 			"online":   true,
		// 		},
		// 	})
		// }
		u.BroadcastMessage(roomCode, entities.WebSocketRequest{
			Method: "system",
			Data: map[string]interface{}{
				"message":  fmt.Sprintf("%s ได้เข้าร่วมห้อง", user.Username),
				"username": user.Username,
				"online":   true,
			},
		})
	}()

	if u.connections[roomCode] == nil {
		u.connections[roomCode] = make(map[string]*websocket.Conn)
	}

	u.connections[roomCode][user.Username] = conn

	log.Printf("Current connections in room %s: %v", roomCode, u.connections[roomCode])

	return nil
}

func (u *websocketUse) RemoveConnection(roomCode string, username string) {
	u.mu.Lock()
	defer func() {
		u.mu.Unlock()
		u.BroadcastMessage(roomCode, entities.WebSocketRequest{
			Method: "system",
			Data: map[string]interface{}{
				"message":  fmt.Sprintf("%s ได้ออกจากห้อง", username),
				"username": username,
				"online":   false,
			},
		})
		fmt.Println("Remove connection for user", username)
		fmt.Println("Current connections in room", roomCode, u.connections[roomCode])
	}()

	if u.connections[roomCode] != nil {
		delete(u.connections[roomCode], username)

		if len(u.connections[roomCode]) == 0 {
			delete(u.connections, roomCode)
		}
	}
}

func (u *websocketUse) BroadcastMessage(roomCode string, req entities.WebSocketRequest) {
	u.mu.RLock()
	room, ok := u.connections[roomCode]
	u.mu.RUnlock()

	if !ok {
		log.Printf("Room %s not found", roomCode)
		return
	}

	log.Printf("Broadcasting message to room %s with %d connections", roomCode, len(room))

	var disconnectedUsers []string

	for username, conn := range room {
		if conn == nil {
			log.Printf("Connection for user %s is nil", username)
			disconnectedUsers = append(disconnectedUsers, username)
			continue
		}

		if err := conn.WriteJSON(req); err != nil {
			log.Printf("Error sending message to %s: %v", username, err)
			disconnectedUsers = append(disconnectedUsers, username)
			conn.Close()
		} else {
			log.Printf("Message sent to %s", username)
		}
	}

	if len(disconnectedUsers) > 0 {
		u.mu.Lock()
		for _, username := range disconnectedUsers {
			delete(u.connections[roomCode], username)
		}
		if len(u.connections[roomCode]) == 0 {
			log.Printf("Room %s is empty, deleting", roomCode)
			delete(u.connections, roomCode)
		}
		u.mu.Unlock()
	}
}

func (u *websocketUse) GetUsersInRoom(roomCode string) []string {
	u.mu.RLock()
	defer u.mu.RUnlock()

	if room, ok := u.connections[roomCode]; ok {
		var users []string
		for username := range room {
			users = append(users, username)
		}
		return users
	}
	return nil
}
