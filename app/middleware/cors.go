package middleware

import (
	"app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORSMiddleware() fiber.Handler {
	return cors.New(cors.Config{
    AllowOrigins: utils.Env["CORS_ALLOWED_ORIGINS"],
    AllowHeaders: "Origin, Content-Type, Accept",
		ExposeHeaders: "Content-Length",
		MaxAge: 43200,  // 12 hours
	})
}
