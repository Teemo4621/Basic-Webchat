package middlewares

import (
	"strings"

	"github.com/Teemo4621/Basic-Webchat/configs"
	"github.com/Teemo4621/Basic-Webchat/pkgs/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	Cfg *configs.Config
}

func NewAuthMiddleware(cfg *configs.Config) *AuthMiddleware {
	return &AuthMiddleware{Cfg: cfg}
}

func (a *AuthMiddleware) JwtAuthentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		if accessToken == "" {
			return utils.UnauthorizedResponse(c, "Unauthorized")
		}

		tokenData, err := utils.ParseAccessToken(a.Cfg, accessToken)
		if err != nil {
			return utils.UnauthorizedResponse(c, "Unauthorized")
		}

		c.Locals("userID", tokenData.Id)

		return c.Next()
	}
}

func (a *AuthMiddleware) WebSocketAuthentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := strings.TrimPrefix(c.Query("token"), "Bearer ")
		if accessToken == "" {
			return utils.UnauthorizedResponse(c, "Unauthorized")
		}

		tokenData, err := utils.ParseAccessToken(a.Cfg, accessToken)
		if err != nil {
			return utils.UnauthorizedResponse(c, "Unauthorized")
		}

		c.Locals("userID", tokenData.Id)
		c.Locals("userUsername", tokenData.Username)

		return c.Next()
	}
}
