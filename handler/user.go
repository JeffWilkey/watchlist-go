package handler

import (
	"fmt"

	"github.com/jeffwilkey/what-to-watch/database"
	"github.com/jeffwilkey/what-to-watch/model"
	"github.com/jeffwilkey/what-to-watch/services"

	"github.com/gofiber/fiber/v2"
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
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request",
		})
	}

	hash, err := services.HashPassword(user.Password)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error hashing password",
		})
	}

	user.Password = hash
	insertionResult, err := collection.InsertOne(c.Context() , &user); 

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Printf("User created with _id: %v\n", insertionResult.InsertedID)

	newUser := NewUser{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}