package service

import (
	"errors"
	"time"

	"github.com/jeffwilkey/watchlist-go/database"
	"github.com/jeffwilkey/watchlist-go/dto"
	"github.com/jeffwilkey/watchlist-go/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindMediaByWatchlistId(c *fiber.Ctx, watchlistId primitive.ObjectID) ([]model.Media, error) {
	// Get all media from database
	collection := database.Mongo.Db.Collection("media")
	cursor, err := collection.Find(c.Context(), bson.M{"watchlistId": watchlistId})
	if err != nil {
		return nil, err
	}

	var results []model.Media
	if err := cursor.All(c.Context(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func FindMediaById(c *fiber.Ctx, id primitive.ObjectID) (*model.Media, error) {
	collection := database.Mongo.Db.Collection("media")
	media := new(model.Media)

	err := collection.FindOne(c.Context(), model.Media{ID: id}).Decode(&media)
	if err != nil {
		return nil, err
	}

	return media, nil
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

func UpdateMedia(c *fiber.Ctx, id primitive.ObjectID, input dto.MediaUpdateRequest, media *model.Media) (int, error) {
	collection := database.Mongo.Db.Collection("media")

	update := bson.M{
		"$set": media,
	}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	update["updatedAt"] = primitive.NewDateTimeFromTime(time.Now().UTC())

	updateResult := collection.FindOneAndUpdate(c.Context(), bson.M{"_id": id}, update, &opt)
	err := updateResult.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.StatusNotFound, errors.New("Media not found")
		}
		return fiber.StatusInternalServerError, err
	}

	err = updateResult.Decode(&media)

	if err != nil {
		return fiber.StatusInternalServerError, err
	}

	return fiber.StatusOK, nil
}

func DeleteMediaByWatchlistId(c *fiber.Ctx, watchlistId primitive.ObjectID) error {
	collection := database.Mongo.Db.Collection("media")

	_, err := collection.DeleteMany(c.Context(), bson.M{"watchlistId": watchlistId})
	if err != nil {
		return err
	}

	return nil
}

func DeleteMedia(c *fiber.Ctx, id primitive.ObjectID) (int, error) {
	collection := database.Mongo.Db.Collection("media")
	_, err := collection.DeleteOne(c.Context(), bson.M{"_id": id})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.StatusNotFound, errors.New("Media not found")
		}
		return fiber.StatusInternalServerError, err
	}
	return fiber.StatusNoContent, nil
}
