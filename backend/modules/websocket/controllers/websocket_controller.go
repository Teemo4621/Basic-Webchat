package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/Teemo4621/Basic-Webchat/modules/entities"
	"github.com/Teemo4621/Basic-Webchat/pkgs/middlewares"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type WebSocketCon struct {
	Cfg              websocket.Config
	WebSocketUsecase entities.WebSocketUsecase
	AuthMiddleware   middlewares.AuthMiddleware
}

func NewWebSocketController(r fiber.Router, cfg websocket.Config, webSocketUsecase entities.WebSocketUsecase, authMiddleware middlewares.AuthMiddleware) {
	controllers := &WebSocketCon{
		Cfg:              cfg,
		WebSocketUsecase: webSocketUsecase,
		AuthMiddleware:   authMiddleware,
	}

	r.Get("/", controllers.AuthMiddleware.WebSocketAuthentication(), websocket.New(controllers.WebSocket, controllers.Cfg))
}

func (c *WebSocketCon) WebSocket(conn *websocket.Conn) {
	userID := conn.Locals("userID").(uint)
	userUsername := conn.Locals("userUsername").(string)

	roomCode := conn.Params("room_code")
	if roomCode == "" {
		conn.WriteJSON(entities.WebSocketResponse{
			Method: "error",
			Data:   "room code is required",
		})
		return
	}

	defer func() {
		conn.Close()
		c.WebSocketUsecase.RemoveConnection(roomCode, userUsername)
	}()

	if err := c.WebSocketUsecase.AddConnection(roomCode, userID, conn); err != nil {
		conn.WriteJSON(entities.WebSocketResponse{
			Method: "error",
			Data:   err.Error(),
		})
		return
	}

	for _, username := range c.WebSocketUsecase.GetUsersInRoom(roomCode) {
		if username == userUsername {
			continue
		}
		conn.WriteJSON(entities.WebSocketResponse{
			Method: "system",
			Data: map[string]interface{}{
				"message":  fmt.Sprintf("%s ได้เข้าร่วมห้อง", username),
				"username": username,
				"online":   true,
			},
		})
	}

	for {
		var req entities.WebSocketRequest
		if err := conn.ReadJSON(&req); err != nil {
			log.Println(err.Error())
			break
		}

		switch req.Method {
		case "message":
			req.Data.(map[string]interface{})["username"] = userUsername
			req.Data.(map[string]interface{})["created_at"] = time.Now()
			c.WebSocketUsecase.BroadcastMessage(roomCode, req)
		}
	}
}
