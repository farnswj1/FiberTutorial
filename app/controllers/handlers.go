package controllers

import "github.com/gofiber/fiber/v2"

func DefaultErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(
		&fiber.Map{"error": err.Error()},
	)
}
