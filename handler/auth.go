package handler

import (
	"github.com/jeffwilkey/what-to-watch/config"
	"github.com/jeffwilkey/what-to-watch/model"
	"github.com/jeffwilkey/what-to-watch/services"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email	  string `json:"email"`
		Password  string `json:"password"`
	}
	type UserData struct {
		ID        primitive.ObjectID `json:"id"`
		FirstName string 			 `json:"firstName"`
		LastName  string             `json:"lastName"`
		Email	  string 			 `json:"email"`
		Password  string 			 `json:"password"`
	}

	input := new(LoginInput)
	var userData UserData

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err })
	}

	email := input.Email
	password := input.Password
	user := new(model.User)
	
	user = services.FindUserByEmail(email)

	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil })
	} else {
		userData = UserData{
			ID: 	   user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email: 	   user.Email,
			Password:  user.Password,
		}
	}

	if !services.CheckPasswordHash(password, userData.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil })
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = userData.Email
	claims["firstName"] = userData.FirstName
	claims["lastName"] = userData.LastName
	claims["userId"] = userData.ID
	
	t, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error creating token", "data": err })
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Logged in successfully", "data": fiber.Map{"token": t}})
}