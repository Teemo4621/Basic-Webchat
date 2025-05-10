package controllers

import (
	"github.com/Teemo4621/Basic-Webchat/configs"
	"github.com/Teemo4621/Basic-Webchat/modules/entities"
	"github.com/Teemo4621/Basic-Webchat/pkgs/middlewares"
	"github.com/Teemo4621/Basic-Webchat/pkgs/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type messageCon struct {
	Cfg            configs.Config
	MessageUsecase entities.MessageUsecase
	AuthMiddleware middlewares.AuthMiddleware
}

func NewMessageController(r fiber.Router, cfg configs.Config, messageUsecase entities.MessageUsecase, authMiddleware middlewares.AuthMiddleware) {
	controllers := &messageCon{Cfg: cfg, MessageUsecase: messageUsecase, AuthMiddleware: authMiddleware}
	r.Get("/", authMiddleware.JwtAuthentication(), controllers.GetMessagesByRoomId)
	r.Post("/", authMiddleware.JwtAuthentication(), controllers.SendMessage)
	r.Post("/:message_id/delete", authMiddleware.JwtAuthentication(), controllers.DeleteMessage)
}

func (a *messageCon) GetMessagesByRoomId(c *fiber.Ctx) error {
	roomCode := c.Params("room_code")
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	if roomCode == "" {
		return utils.BadRequestResponse(c, "BadRequest")
	}

	messages, err := a.MessageUsecase.GetMessagesByRoomId(roomCode, page, limit)
	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	return utils.OkResponse(c, fiber.Map{
		"messages": messages.Messages,
		"pagination": fiber.Map{
			"page":       page,
			"limit":      limit,
			"page_total": messages.PageTotal,
		},
	})
}

func (a *messageCon) SendMessage(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uint)
	roomCode := c.Params("room_code")

	var req entities.MessageRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequestResponse(c, "BadRequest")
	}

	if req.Content == "" {
		return utils.BadRequestResponse(c, "BadRequest")
	}

	message, err := a.MessageUsecase.SendMessage(roomCode, userId, req.Content)
	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	return utils.OkResponse(c, message)
}

func (a *messageCon) DeleteMessage(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uint)
	roomCode := c.Params("room_code")
	messageId := c.Params("message_id")

	messageUUID, err := uuid.Parse(messageId)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid message ID")
	}

	if err := a.MessageUsecase.DeleteMessage(roomCode, userId, messageUUID); err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	return utils.OkResponse(c, "delete message successfully!")
}
