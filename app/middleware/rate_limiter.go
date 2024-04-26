package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func KeyGeneratorHandler(c *fiber.Ctx) string {
	return c.IP()
}

func LimitReachedHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusTooManyRequests).JSON(
		&fiber.Map{"error": "Too many requests"},
	)
}

func RateLimiterMiddleware(limit int, duration int) fiber.Handler {
	return limiter.New(limiter.Config{
		Expiration: time.Duration(duration) * time.Second,
		Max: limit,
		KeyGenerator: KeyGeneratorHandler,
		LimitReached: LimitReachedHandler,
	})
}
