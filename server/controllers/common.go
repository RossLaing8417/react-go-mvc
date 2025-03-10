package controllers

import "github.com/gofiber/fiber/v2"

func respondError(c *fiber.Ctx, code int, err error, message string) error {
	return c.Status(code).JSON(fiber.Map{
		"error":   err.Error(),
		"message": message,
	})
}
