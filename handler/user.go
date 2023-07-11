package handler

import (
	"github.com/jeffwilkey/watchlist-go/dto"
	"github.com/jeffwilkey/watchlist-go/model"
	"github.com/jeffwilkey/watchlist-go/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(&user); err != nil {
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
	var input dto.UserUpdateRequest

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
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	token := c.Locals("user").(*jwt.Token)

	// Validate JWT token
	if !service.ValidToken(token, id) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Unauthorized", "data": nil})
	}

	// Update user in database
	var user model.User
	status, err := service.UpdateUser(c, id, input, &user)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User updated", "data": dto.CreateUserResponse(user)})
}

func DeleteUser(c *fiber.Ctx) error {
	// Format user ID from param and get token
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	token := c.Locals("user").(*jwt.Token)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid ID", "data": err})
	}

	// Validate JWT token
	if !service.ValidToken(token, id) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Unauthorized", "data": nil})
	}

	// Delete user from database
	status, err := service.DeleteUser(c, id)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": err})
	}

	return c.Status(status).JSON(fiber.Map{"status": "success", "message": "User deleted", "data": nil})
}
