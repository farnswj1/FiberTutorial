package controllers

import (
	"app/validators"

	"github.com/gofiber/fiber/v2"
)

func Root(c *fiber.Ctx) error {
	return c.Render("index", &fiber.Map{"IP": c.IP()})
}

func Greet(c *fiber.Ctx) error {
	name := c.Params("name", "user")
	return c.JSON(&fiber.Map{
		"greeting": "Hello, " + name + "!",
	})
}

func Add(c *fiber.Ctx) error {
	var body *validators.Number

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{"error": "Unable to parse data"},
		)
	}

	if err := validators.Validator.Struct(body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{"error": err.Error()},
		)
	}

	return c.JSON(body.A + body.B)
}
