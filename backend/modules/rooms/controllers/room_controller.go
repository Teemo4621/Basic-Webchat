package controllers

import (
	"github.com/Teemo4621/Basic-Webchat/configs"
	"github.com/Teemo4621/Basic-Webchat/modules/entities"
	"github.com/Teemo4621/Basic-Webchat/pkgs/middlewares"
	"github.com/Teemo4621/Basic-Webchat/pkgs/utils"
	"github.com/gofiber/fiber/v2"
)

type roomCon struct {
	Cfg            configs.Config
	RoomUsecase    entities.RoomUsecase
	AuthMiddleware middlewares.AuthMiddleware
}

func NewRoomController(r fiber.Router, cfg configs.Config, roomUsecase entities.RoomUsecase, authMiddleware middlewares.AuthMiddleware) {
	controllers := &roomCon{Cfg: cfg, RoomUsecase: roomUsecase, AuthMiddleware: authMiddleware}
	r.Get("/", authMiddleware.JwtAuthentication(), controllers.GetRoomsByUserId)
	r.Post("/", authMiddleware.JwtAuthentication(), controllers.CreateRoom)
	r.Get("/:room_code", authMiddleware.JwtAuthentication(), controllers.GetRoom)
	r.Get("/:room_code/members", authMiddleware.JwtAuthentication(), controllers.GetRoomMembers)
	r.Post("/:room_code/join", authMiddleware.JwtAuthentication(), controllers.JoinRoom)
	r.Post("/:room_code/leave", authMiddleware.JwtAuthentication(), controllers.LeaveRoom)
	r.Post("/:room_code/delete", authMiddleware.JwtAuthentication(), controllers.DeleteRoom)
}

func (a *roomCon) GetRoomsByUserId(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uint)
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	rooms, err := a.RoomUsecase.GetRoomsByUserId(userId, page, limit)
	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	return utils.OkResponse(c, fiber.Map{
		"rooms": rooms.Rooms,
		"pagination": fiber.Map{
			"page":       page,
			"limit":      limit,
			"page_total": rooms.PageTotal,
		},
	})
}

func (a *roomCon) GetRoom(c *fiber.Ctx) error {
	roomCode := c.Params("room_code")

	room, err := a.RoomUsecase.GetRoom(roomCode)
	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	return utils.OkResponse(c, room)
}

func (a *roomCon) CreateRoom(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uint)

	var req entities.RoomCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequestResponse(c, "BadRequest")
	}

	if req.Name == "" || req.Description == "" {
		return utils.BadRequestResponse(c, "BadRequest")
	}

	room := entities.Room{
		OwnerID:     userId,
		Name:        req.Name,
		Description: req.Description,
	}

	newRoom, err := a.RoomUsecase.CreateRoom(&room)
	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	return utils.OkResponse(c, newRoom)
}

func (a *roomCon) JoinRoom(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uint)
	roomCode := c.Params("room_code")

	if roomCode == "" {
		return utils.BadRequestResponse(c, "BadRequest")
	}

	if err := a.RoomUsecase.JoinRoom(roomCode, userId); err != nil {
		return utils.NotFoundResponse(c, "Room not found")
	}

	return utils.OkResponse(c, "join room successfully!")
}

func (a *roomCon) LeaveRoom(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uint)
	roomCode := c.Params("room_code")

	if roomCode == "" {
		return utils.BadRequestResponse(c, "BadRequest")
	}

	if err := a.RoomUsecase.LeaveRoom(roomCode, userId); err != nil {
		return utils.NotFoundResponse(c, "Room not found")
	}

	return utils.OkResponse(c, "leave room successfully!")
}

func (a *roomCon) DeleteRoom(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uint)
	roomCode := c.Params("room_code")

	if roomCode == "" {
		return utils.BadRequestResponse(c, "BadRequest")
	}

	if err := a.RoomUsecase.DeleteRoom(roomCode, userId); err != nil {
		return utils.NotFoundResponse(c, "Room not found")
	}

	return utils.OkResponse(c, "delete room successfully!")
}

func (a *roomCon) GetRoomMembers(c *fiber.Ctx) error {
	roomCode := c.Params("room_code")

	if roomCode == "" {
		return utils.BadRequestResponse(c, "BadRequest")
	}

	members, err := a.RoomUsecase.GetRoomMembers(roomCode)
	if err != nil {
		return utils.NotFoundResponse(c, "Room not found")
	}

	return utils.OkResponse(c, members)
}