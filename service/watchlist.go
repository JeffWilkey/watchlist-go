package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeffwilkey/watchlist-go/database"
	"github.com/jeffwilkey/watchlist-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindWatchlistsByOwnerId(c *fiber.Ctx, ownerId primitive.ObjectID) ([]model.Watchlist, error) {
	watchlists := make([]model.Watchlist, 0)

	collection := database.Mongo.Db.Collection("watchlists")
	cursor, err := collection.Find(c.Context(), bson.M{"ownerId": ownerId})
	if err != nil {
		return watchlists, err
	}
	err = cursor.All(c.Context(), &watchlists)
	if err != nil {
		return watchlists, err
	}
	return watchlists, nil
}
