package router

import (
	"github.com/jeffwilkey/watchlist-go/handler"
	"github.com/jeffwilkey/watchlist-go/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// User
	user := api.Group("/users")
	user.Post("/", handler.CreateUser)
	user.Put("/:id", middleware.Protected(), handler.UpdateUser)
}