package services

import (
	"context"

	"github.com/jeffwilkey/what-to-watch/database"
	"github.com/jeffwilkey/what-to-watch/model"

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

func FindUserByEmail(email string) (*model.User) {
	collection := database.Mongo.Db.Collection("users")

	user := new(model.User)

	collection.FindOne(context.TODO(), model.User{Email: email}).Decode(&user)

	return user
}