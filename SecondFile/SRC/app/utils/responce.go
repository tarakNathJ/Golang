package utils

import "github.com/gofiber/fiber/v2"

func JSONResponce(c *fiber.Ctx, status int, data interface{}, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"success": true,
		"message": message,
		"data":    data,
	})
}
