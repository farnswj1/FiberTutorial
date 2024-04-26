package main

import (
	"app/controllers"
	"app/middleware"
	"app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func GetApp() *fiber.App {
	utils.LoadEnv()

	app := fiber.New(fiber.Config{
		Views: html.New("templates", ".html"),
	})

	// Middleware
	app.Use(middleware.RateLimiterMiddleware(5, 60))
	app.Use(middleware.RecoverMiddleware())
	app.Use(middleware.LoggerMiddleware())
	app.Use(middleware.CORSMiddleware())
	app.Use(middleware.HelmetMiddleware())
	app.Use(middleware.IdempotencyMiddleware())

	// Routes
	app.Get("/", controllers.Root)
	app.Get("/users/:name", controllers.Greet)
	app.Post("/add", controllers.Add)

	return app
}

func main() {
	app := GetApp()
	app.Listen(":8080")
}
