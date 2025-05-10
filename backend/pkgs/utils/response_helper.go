package utils

import "github.com/gofiber/fiber/v2"

func OkResponse(r *fiber.Ctx, data interface{}) error {
	return r.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    data,
		"status":  "ok",
		"message": "success",
	})
}

func NotFoundResponse(r *fiber.Ctx, message string) error {
	return r.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": message,
		"status":  "error",
	})
}

func ErrorResponse(r *fiber.Ctx, message string) error {
	return r.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": message,
		"status":  "error",
	})
}

func BadRequestResponse(r *fiber.Ctx, message string) error {
	return r.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": message,
		"status":  "error",
	})
}

func UnauthorizedResponse(r *fiber.Ctx, message string) error {
	return r.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": message,
		"status":  "error",
	})
}
