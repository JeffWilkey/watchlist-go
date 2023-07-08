package handler

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jeffwilkey/watchlist-go/database"
	"github.com/jeffwilkey/watchlist-go/model"
	"github.com/jeffwilkey/watchlist-go/service"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateUser new user
func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		FirstName	string	`json:"firstName"`
		LastName	string	`json:"lastName"`
		Email		string	`json:"email"`
	}

	collection := database.Mongo.Db.Collection("users")
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request", "data": err})
	}

	hash, err := service.HashPassword(user.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error hashing password", "data": err})
	}

	user.Password = hash
	insertionResult, err := collection.InsertOne(c.Context() , &user); 

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err.Error()})
	}

	fmt.Printf("User created with _id: %v\n", insertionResult.InsertedID)

	newUser := NewUser{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

func UpdateUser(c *fiber.Ctx) error {
	type UpdateUserInput struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	var input UpdateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid input", "data": err})
	}

	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !service.ValidToken(token, id) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Unauthorized", "data": nil})
	}

	collection := database.Mongo.Db.Collection("users")
	var user model.User

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "firstName", Value: input.FirstName},
		{Key: "lastName", Value: input.LastName},
	}}}

	updateResult := collection.FindOneAndUpdate(c.Context(), filter, update)
	err := updateResult.Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't update user", "data": err})
	}

	err = updateResult.Decode(&user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't decode user", "data": err})
	}
	
	return c.JSON(fiber.Map{"status": "success", "message": "User updated", "data": user})
}	