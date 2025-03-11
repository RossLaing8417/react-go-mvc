package controllers

import "github.com/gofiber/fiber/v2"

func responseError(c *fiber.Ctx, status int, err error, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"error":   err.Error(),
		"message": message,
	})
}
