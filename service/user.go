package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeffwilkey/watchlist-go/database"
	"github.com/jeffwilkey/watchlist-go/model"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func FindUserByEmail(c *fiber.Ctx, email string) *model.User {
	collection := database.Mongo.Db.Collection("users")

	user := new(model.User)

	collection.FindOne(c.Context(), model.User{Email: email}).Decode(&user)

	return user
}

func ValidToken(t *jwt.Token, id string) bool {
	claims := t.Claims.(jwt.MapClaims)

	return claims["userId"] == id
}
