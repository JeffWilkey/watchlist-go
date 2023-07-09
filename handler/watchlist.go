package handler

import (
	"github.com/jeffwilkey/watchlist-go/database"
	"github.com/jeffwilkey/watchlist-go/dto"
	"github.com/jeffwilkey/watchlist-go/model"
	"github.com/jeffwilkey/watchlist-go/service"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	return c.JSON(fiber.Map{"status": "success", "message": "All watchlists", "data": dto.CreateWatchlistsResponse(results)})
}

func GetWatchlist(c *fiber.Ctx) error {
	// Format watchlist ID from param and get token
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Missing watchlist ID"})
	}

	// Get watchlist from database
	collection := database.Mongo.Db.Collection("watchlists")
	var watchlist model.Watchlist

	filter := bson.D{{Key: "_id", Value: id}}
	err = collection.FindOne(c.Context(), filter).Decode(&watchlist)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Watchlist not found", "data": nil})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't get watchlist", "data": err.Error()})
	}

	media, err := service.GetMediaByWatchlistId(c, watchlist.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't get media", "data": err.Error()})
	}

	mediaList := dto.CreateMediaListResponse(media)

	return c.JSON(fiber.Map{"status": "success", "message": "Watchlist found", "data": dto.CreateWatchlistWithMediaResponse(watchlist, mediaList)})
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
