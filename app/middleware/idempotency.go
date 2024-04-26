package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
)

func IdempotencyMiddleware() fiber.Handler {
	return idempotency.New()
}
