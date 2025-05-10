package servers

import (
	_authHttp "github.com/Teemo4621/Basic-Webchat/modules/auth/controllers"
	_authRepository "github.com/Teemo4621/Basic-Webchat/modules/auth/repositories"
	_authUsecase "github.com/Teemo4621/Basic-Webchat/modules/auth/usecases"
	_messageHttp "github.com/Teemo4621/Basic-Webchat/modules/messages/controllers"
	_messageRepository "github.com/Teemo4621/Basic-Webchat/modules/messages/repositories"
	_messageUsecase "github.com/Teemo4621/Basic-Webchat/modules/messages/usecases"
	_roomMembersRepository "github.com/Teemo4621/Basic-Webchat/modules/roommembers/repositories"
	_roomHttp "github.com/Teemo4621/Basic-Webchat/modules/rooms/controllers"
	_roomRepository "github.com/Teemo4621/Basic-Webchat/modules/rooms/repositories"
	_roomUsecase "github.com/Teemo4621/Basic-Webchat/modules/rooms/usecases"
	_userRepository "github.com/Teemo4621/Basic-Webchat/modules/users/repositories"
	_websocketHttp "github.com/Teemo4621/Basic-Webchat/modules/websocket/controllers"
	_websocketUsecase "github.com/Teemo4621/Basic-Webchat/modules/websocket/usecases"
	"github.com/Teemo4621/Basic-Webchat/pkgs/middlewares"
	"github.com/Teemo4621/Basic-Webchat/pkgs/utils"
	"github.com/gofiber/contrib/websocket"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (s *Server) MapHandlers() error {
	s.App.Use(cors.New())

	//middlewares
	authReq := middlewares.NewAuthMiddleware(s.Cfg)
	loggerMiddleware := middlewares.LoggerMiddleware()

	//repositories
	authRepository := _authRepository.NewAuthRepository(s.Db)
	messageRepository := _messageRepository.NewMessageRepository(s.Db)
	roomRepository := _roomRepository.NewRoomRepository(s.Db)
	roomMemberRepository := _roomMembersRepository.NewRoomMemberRepository(s.Db)

	// Group a api
	apiGroup := s.App.Group("/api", loggerMiddleware)
	v1 := apiGroup.Group("/v1")
	authGroup := v1.Group("/auth")
	roomGroup := v1.Group("/rooms")
	messageGroup := roomGroup.Group("/:room_code/messages")
	websocketGroup := roomGroup.Group("/:room_code/ws")

	// User group
	userRepository := _userRepository.NewUserRepository(s.Db)

	// Auth group
	authUsecase := _authUsecase.NewAuthUsecase(authRepository, userRepository)
	_authHttp.NewAuthController(authGroup, *s.Cfg, authUsecase, *authReq)

	// Message group
	messageUsecase := _messageUsecase.NewMessageUsecase(messageRepository, roomRepository, roomMemberRepository)
	_messageHttp.NewMessageController(messageGroup, *s.Cfg, messageUsecase, *authReq)

	// Room group
	roomUsecase := _roomUsecase.NewRoomUsecase(roomRepository, roomMemberRepository, messageRepository, userRepository)
	_roomHttp.NewRoomController(roomGroup, *s.Cfg, roomUsecase, *authReq)

	// WebSocket group
	websocketUseCase := _websocketUsecase.NewWebSocketUsecase(userRepository, roomRepository, roomMemberRepository)
	websocketCfg := websocket.Config{
		RecoverHandler: func(conn *websocket.Conn) {
			if err := recover(); err != nil {
				conn.WriteJSON(fiber.Map{"error": "unexpected error occurred"})
				conn.Close()
			}
		},
	}

	_websocketHttp.NewWebSocketController(websocketGroup, websocketCfg, websocketUseCase, *authReq)

	// End point not found response
	s.App.Use(func(c *fiber.Ctx) error {
		return utils.ErrorResponse(c, "end point not found")
	})

	return nil
}
