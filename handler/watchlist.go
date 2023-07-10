package handler

import (
	"github.com/jeffwilkey/watchlist-go/dto"
	"github.com/jeffwilkey/watchlist-go/service"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetWatchlists(c *fiber.Ctx) error {
	// Get ownerId from query parameters
	ownerId, err := primitive.ObjectIDFromHex(c.Query("ownerId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Missing ownerId"})
	}

	// Get all watchlists from database
	results, err := service.FindWatchlistsByOwnerId(c, ownerId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't get watchlists", "data": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "All watchlists", "data": dto.CreateWatchlistsResponse(results)})
}

func CreateWatchlist(c *fiber.Ctx) error {
	return c.SendString("Create watchlist")
}

func UpdateWatchlist(c *fiber.Ctx) error {
	return c.SendString("Update watchlist")
}

func DeleteWatchlist(c *fiber.Ctx) error {
	return c.SendString("Delete watchlist")
}
