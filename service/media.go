package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeffwilkey/watchlist-go/database"
	"github.com/jeffwilkey/watchlist-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMediaByWatchlistId(c *fiber.Ctx, watchlistId primitive.ObjectID) ([]model.Media, error) {
	// Get all media from database
	collection := database.Mongo.Db.Collection("media")
	cursor, err := collection.Find(c.Context(), bson.D{{Key: "watchlistId", Value: watchlistId}})
	if err != nil {
		return nil, err
	}

	var results []model.Media
	if err := cursor.All(c.Context(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
