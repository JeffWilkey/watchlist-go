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
	user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Watchlist
	watchlist := api.Group("/watchlists")
	watchlist.Get("/", middleware.Protected(), handler.GetWatchlists)
	watchlist.Post("/", middleware.Protected(), handler.CreateWatchlist)
	watchlist.Patch("/:id", middleware.Protected(), handler.UpdateWatchlist)

	// Media
	media := api.Group("/media")
	media.Get("/", middleware.Protected(), handler.GetMediaList)
}
