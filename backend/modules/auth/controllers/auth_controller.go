package controllers

import (
	"log"

	"github.com/Teemo4621/Basic-Webchat/configs"
	"github.com/Teemo4621/Basic-Webchat/modules/entities"
	"github.com/Teemo4621/Basic-Webchat/pkgs/middlewares"
	"github.com/Teemo4621/Basic-Webchat/pkgs/utils"
	"github.com/gofiber/fiber/v2"
)

type authCon struct {
	Cfg            configs.Config
	AuthUsecase    entities.AuthUsecase
	AuthMiddleware middlewares.AuthMiddleware
}

func NewAuthController(r fiber.Router, cfg configs.Config, authUsecase entities.AuthUsecase, authMiddleware middlewares.AuthMiddleware) {
	controller := &authCon{Cfg: cfg, AuthUsecase: authUsecase, AuthMiddleware: authMiddleware}
	r.Post("/login", controller.Login)
	r.Post("/register", controller.Register)
	r.Get("/@me", authMiddleware.JwtAuthentication(), controller.Me)
	r.Post("/refresh-token", controller.RefreshToken)
}

func (a *authCon) Login(c *fiber.Ctx) error {
	var req entities.AuthLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequestResponse(c, "BadRequest")
	}

	// validate body
	if req.Username == "" || req.Password == "" {
		return utils.BadRequestResponse(c, "BadRequest")
	}

	user, err := a.AuthUsecase.Login(&a.Cfg, &req)
	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	return utils.OkResponse(c, user)
}

func (a *authCon) Register(c *fiber.Ctx) error {
	var req entities.AuthRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequestResponse(c, "BadRequest")
	}

	// validate body
	if req.Username == "" || req.Password == "" {
		return utils.BadRequestResponse(c, "BadRequest")
	}

	_, err := a.AuthUsecase.Register(&req)
	if err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	return utils.OkResponse(c, "register successfully!")
}

func (a *authCon) Me(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uint)

	user, err := a.AuthUsecase.Me(&a.Cfg, userId)
	if err != nil {
		return utils.UnauthorizedResponse(c, "Unauthorized")
	}

	return utils.OkResponse(c, user)
}

func (a *authCon) RefreshToken(c *fiber.Ctx) error {
	var req entities.AuthRefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequestResponse(c, "BadRequest")
	}

	// validate body
	if req.RefreshToken == "" {
		return utils.BadRequestResponse(c, "BadRequest")
	}

	user, err := a.AuthUsecase.RefreshToken(&a.Cfg, &req)
	if err != nil {
		log.Println(err.Error())
		return utils.NotFoundResponse(c, "RefreshToken Error")
	}

	return utils.OkResponse(c, user)
}
