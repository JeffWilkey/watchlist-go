package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Watchlist struct {
	ID            primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string               `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=1,max=32"`
	Description   string               `json:"description,omitempty" bson:"description,omitempty" validate:"max=256"`
	OwnerID       primitive.ObjectID   `json:"ownerId,omitempty" bson:"ownerId,omitempty" validate:"required"`
	Collaborators []primitive.ObjectID `json:"collaborators,omitempty" bson:"collaborators,omitempty"`
	CreatedAt     primitive.DateTime   `json:"createdAt,omitempty" bson:"createdAt,omitempty" validate:"required"`
	UpdatedAt     primitive.DateTime   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" validate:"required"`
}
