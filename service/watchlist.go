package service

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jeffwilkey/watchlist-go/database"
	"github.com/jeffwilkey/watchlist-go/dto"
	"github.com/jeffwilkey/watchlist-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindWatchlistById(c *fiber.Ctx, id primitive.ObjectID) (*model.Watchlist, error) {
	collection := database.Mongo.Db.Collection("watchlists")
	watchlist := new(model.Watchlist)

	err := collection.FindOne(c.Context(), model.Watchlist{ID: id}).Decode(&watchlist)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("Watchlist not found")
		}
		return nil, err
	}
	return watchlist, nil
}

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

func CreateWatchlist(c *fiber.Ctx, watchlist *model.Watchlist) error {
	collection := database.Mongo.Db.Collection("watchlists")
	insertResult, err := collection.InsertOne(c.Context(), watchlist)
	if err != nil {
		return err
	}
	watchlist.ID = insertResult.InsertedID.(primitive.ObjectID)
	return nil
}

func UpdateWatchlist(c *fiber.Ctx, id primitive.ObjectID, input dto.WatchlistUpdateRequest, watchlist *model.Watchlist) (int, error) {
	collection := database.Mongo.Db.Collection("watchlists")

	update := bson.M{
		"$set": input,
	}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	updateResult := collection.FindOneAndUpdate(c.Context(), bson.M{"_id": id}, update, &opt)
	err := updateResult.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.StatusNotFound, errors.New("Watchlist not found")
		}
		return fiber.StatusInternalServerError, err
	}

	err = updateResult.Decode(&watchlist)

	if err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func DeleteWatchlist(c *fiber.Ctx, id primitive.ObjectID) (int, error) {
	collection := database.Mongo.Db.Collection("watchlists")
	_, err := collection.DeleteOne(c.Context(), bson.M{"_id": id})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.StatusNotFound, errors.New("Watchlist not found")
		}
		return fiber.StatusInternalServerError, err
	}
	return fiber.StatusNoContent, nil
}
