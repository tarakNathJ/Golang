package utils

import "github.com/gofiber/fiber/v2"

func ErrorResponce(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"error":   message,
		"success": false,
	})
}
