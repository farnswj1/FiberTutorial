package middleware

import (
	"app/utils"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func errorHandler(c *fiber.Ctx, err error) error {
	errorMessage := err.Error()

	if errorMessage == jwtware.ErrJWTMissingOrMalformed.Error() {
		return c.Status(fiber.StatusBadRequest).JSON(
			&fiber.Map{"error": errorMessage},
		)
	}

	return c.Status(fiber.StatusUnauthorized).JSON(
		&fiber.Map{"error": "invalid or expired JWT"},
	)
}

func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(utils.Env["SECRET_KEY"])},
		ErrorHandler: errorHandler,
	})
}
