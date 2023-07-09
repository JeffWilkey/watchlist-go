package handler

import (
	"github.com/jeffwilkey/watchlist-go/database"
	"github.com/jeffwilkey/watchlist-go/dto"
	"github.com/jeffwilkey/watchlist-go/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetWatchlists(c *fiber.Ctx) error {
	// Get userId from query parameters
	userId, err := primitive.ObjectIDFromHex(c.Query("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Missing userId"})
	}

	// Get all watchlists from database
	collection := database.Mongo.Db.Collection("watchlists")
	cursor, err := collection.Find(c.Context(), bson.D{{Key: "userId", Value: userId}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't get watchlists", "data": err.Error()})
	}

	var results []model.Watchlist
	if err := cursor.All(c.Context(), &results); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't get watchlists", "data": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "All watchlists", "data": dto.CreateWatchListsResponse(results)})
}

func GetWatchlist(c *fiber.Ctx) error {
	return c.SendString("Single watchlist")
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
