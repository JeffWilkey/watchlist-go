package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jeffwilkey/watchlist-go/database"
	"github.com/jeffwilkey/watchlist-go/dto"
	"github.com/jeffwilkey/watchlist-go/model"
	"github.com/jeffwilkey/watchlist-go/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request", "data": err})
	}

	// Validate user input
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Validation failed", "data": err})
	}

	// Create user in database w/ hashed password
	err = service.CreateUser(c, user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error Creating User", "data": err})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Created user", "data": dto.CreateUserResponse(*user)})
}

func UpdateUser(c *fiber.Ctx) error {
	type UpdateUserInput struct {
		FirstName string `json:"firstName" validate:"required,min=1,max=32"`
		LastName  string `json:"lastName" validate:"required,min=2,max=32"`
	}

	var input UpdateUserInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid input", "data": err})
	}

	// Validate user input
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Validation failed", "data": err.Error()})
	}

	// Format user ID from param and get token
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	token := c.Locals("user").(*jwt.Token)

	// Validate JWT token
	if !service.ValidToken(token, idParam) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Unauthorized", "data": nil})
	}

	// Update user in database
	collection := database.Mongo.Db.Collection("users")
	var user model.User

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "firstName", Value: input.FirstName},
		{Key: "lastName", Value: input.LastName},
	}}}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	updateResult := collection.FindOneAndUpdate(c.Context(), filter, update, &opt)
	err = updateResult.Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't update user", "data": err})
	}

	// Decode updated user
	err = updateResult.Decode(&user)

	// Return user DTO
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't decode user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User updated", "data": dto.CreateUserResponse(user)})
}

func DeleteUser(c *fiber.Ctx) error {
	// Format user ID from param and get token
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	token := c.Locals("user").(*jwt.Token)

	// Validate JWT token
	if !service.ValidToken(token, idParam) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Unauthorized", "data": nil})
	}

	// Delete user from database
	collection := database.Mongo.Db.Collection("users")

	filter := bson.D{{Key: "_id", Value: id}}
	deleteResult, err := collection.DeleteOne(c.Context(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't delete user", "data": err})
	}

	if deleteResult.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"status": "success", "message": "User deleted", "data": nil})
}
