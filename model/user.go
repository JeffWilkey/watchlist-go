package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName,omitempty" bson:"firstName,omitempty" validate:"required,min=1,max=32"`
	LastName  string             `json:"lastName,omitempty" bson:"lastName,omitempty" validate:"required,min=2,max=32"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty" validate:"required,min=8,max=100"`
	CreatedAt primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty" validate:"required"`
	UpdatedAt primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" validate:"required"`
}
