package entities

import "github.com/gofiber/contrib/websocket"

type (
	WebSocketUsecase interface {
		AddConnection(roomCode string, userID uint, conn *websocket.Conn) error
		RemoveConnection(roomCode string, username string)
		BroadcastMessage(roomCode string, message WebSocketRequest)
		GetUsersInRoom(roomCode string) []string
	}

	WebSocketRequest struct {
		Method string      `json:"method"`
		Data   interface{} `json:"data"`
	}

	WebSocketResponse struct {
		Method string      `json:"method"`
		Data   interface{} `json:"data"`
	}
)
