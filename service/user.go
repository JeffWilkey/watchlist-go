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

var collection = database.Mongo.Db.Collection("users")

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func FindUserByEmail(c *fiber.Ctx, email string) *model.User {
	user := new(model.User)

	collection.FindOne(c.Context(), model.User{Email: email}).Decode(&user)

	return user
}

func CreateUser(c *fiber.Ctx, user *model.User) error {
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
	// Build update query
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "firstName", Value: input.FirstName},
		{Key: "lastName", Value: input.LastName},
	}}}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	// Update user in database
	updateResult := collection.FindOneAndUpdate(c.Context(), filter, update, &opt)
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

func ValidToken(t *jwt.Token, id string) bool {
	claims := t.Claims.(jwt.MapClaims)

	return claims["userId"] == id
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
