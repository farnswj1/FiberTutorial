package middleware

import (
	"app/cache"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func keyGeneratorHandler(c *fiber.Ctx) string {
	return c.IP()
}

func limitReachedHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusTooManyRequests).JSON(
		&fiber.Map{"error": "Too many requests"},
	)
}

func RateLimiterMiddleware(limit, duration int) fiber.Handler {
	return limiter.New(limiter.Config{
		Expiration: time.Duration(duration) * time.Second,
		Max: limit,
		KeyGenerator: keyGeneratorHandler,
		LimitReached: limitReachedHandler,
		Storage: cache.Redis,
	})
}
