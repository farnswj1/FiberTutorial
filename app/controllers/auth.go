package controllers

import (
	"app/utils"
	"app/validators"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {
	var body *validators.User

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

	if body.Username != "john" || body.Password != "doe" {
		return c.Status(fiber.StatusUnauthorized).JSON(
			&fiber.Map{"error": "Invalid credentials"},
		)
	}

	claims := jwt.MapClaims{
		"id":  1,
		"username": body.Username,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access, err := token.SignedString([]byte(utils.Env["SECRET_KEY"]))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			&fiber.Map{"error": err.Error()},
		)
	}

	return c.JSON(&fiber.Map{"access": access})
}


func Restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["username"].(string)
	return c.JSON(&fiber.Map{"status": "Welcome " + name + "!"})
}
