package services

import (
	"context"

	"github.com/jeffwilkey/what-to-watch/database"
	"github.com/jeffwilkey/what-to-watch/model"
)

// Find a user by email
func FindUserByEmail(email string) (*model.User) {
	collection := database.Mongo.Db.Collection("users")

	user := new(model.User)

	collection.FindOne(context.TODO(), model.User{Email: email}).Decode(&user)

	return user
}