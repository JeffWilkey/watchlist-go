package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeffwilkey/watchlist-go/database"
	"github.com/jeffwilkey/watchlist-go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

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

func ValidToken(t *jwt.Token, id string) bool {
	claims := t.Claims.(jwt.MapClaims)

	return claims["userId"] == id
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
