package router

import (
	"github.com/jeffwilkey/what-to-watch/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())

	// User
	user := api.Group("/users")
	user.Post("/", handler.CreateUser)
}