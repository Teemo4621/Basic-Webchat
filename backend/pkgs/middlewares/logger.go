package middlewares

import (
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		duration := time.Since(start)
		method := color.New(color.FgCyan, color.Bold).Sprint(c.Method())
		url := color.New(color.FgGreen).Sprint(c.OriginalURL())
		query := color.New(color.FgHiMagenta).Sprint(c.Context().QueryArgs())

		log.Printf("📥 %s | %s | Query: %s | ⏱️ %v ",
			method,
			url,
			query,
			duration,
		)

		return c.Next()
	}
}
