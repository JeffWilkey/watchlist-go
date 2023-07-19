package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Media struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty" validate:"required,min=1,max=32"`
	Poster      string             `json:"poster,omitempty" bson:"poster,omitempty" validate:"required,min=2,max=32"`
	ReleaseDate primitive.DateTime `json:"releaseDate,omitempty" bson:"releaseDate,omitempty" validate:"required"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty" validate:"oneof=watched unwatched in-progress"`
	TmdbID      int                `json:"tmdbId,omitempty" bson:"tmdbId,omitempty" validate:"required"`
	WatchlistID primitive.ObjectID `json:"watchlistId,omitempty" bson:"watchlistId,omitempty" validate:"required"`
	CreatedAt   primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty" validate:"required"`
	UpdatedAt   primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" validate:"required"`
}
