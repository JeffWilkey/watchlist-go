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

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidToken(t *jwt.Token, id primitive.ObjectID) bool {
	claims := t.Claims.(jwt.MapClaims)

	return claims["userId"] == id.Hex()
}

func FindUserByEmail(c *fiber.Ctx, email string) *model.User {
	collection := database.Mongo.Db.Collection("users")
	user := new(model.User)

	collection.FindOne(c.Context(), model.User{Email: email}).Decode(&user)

	return user
}

func CreateUser(c *fiber.Ctx, user *model.User) error {
	collection := database.Mongo.Db.Collection("users")
	hash, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash
	insertResult, err := collection.InsertOne(c.Context(), &user)
	if err != nil {
		return err
	}

	user.ID = insertResult.InsertedID.(primitive.ObjectID)

	return nil
}

func UpdateUser(c *fiber.Ctx, id primitive.ObjectID, input dto.UserUpdateRequest, user *model.User) (int, error) {
	collection := database.Mongo.Db.Collection("users")
	// Build update query
	update := bson.M{"$set": input}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	// Update user in database
	updateResult := collection.FindOneAndUpdate(c.Context(), bson.M{"_id": id}, update, &opt)
	err := updateResult.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.StatusNotFound, errors.New("User not found")
		}
		return fiber.StatusInternalServerError, errors.New("Couldn't update user")
	}

	// Decode updated user
	err = updateResult.Decode(&user)

	if err != nil {
		return fiber.StatusInternalServerError, errors.New("Couldn't decode user")
	}

	return fiber.StatusOK, nil
}

func DeleteUser(c *fiber.Ctx, id primitive.ObjectID) (int, error) {
	collection := database.Mongo.Db.Collection("users")
	_, err := collection.DeleteOne(c.Context(), bson.M{"_id": id})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.StatusNotFound, errors.New("User not found")
		}
		return fiber.StatusInternalServerError, err
	}
	return fiber.StatusNoContent, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
