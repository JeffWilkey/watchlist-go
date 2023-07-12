package service

import (
	"time"

	"github.com/jeffwilkey/watchlist-go/database"
	"github.com/jeffwilkey/watchlist-go/model"

	"github.com/gofiber/fiber/v2"
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

func CreateMedia(c *fiber.Ctx, media *model.Media) error {
	collection := database.Mongo.Db.Collection("media")

	media.CreatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
	media.UpdatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())

	insertResult, err := collection.InsertOne(c.Context(), media)
	if err != nil {
		return err
	}

	media.ID = insertResult.InsertedID.(primitive.ObjectID)

	return nil
}
