package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Media struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty" validate:"required,min=1,max=32"`
	Year        string             `json:"year,omitempty" bson:"year,omitempty" validate:"required,min=2,max=32"`
	Poster      string             `json:"poster,omitempty" bson:"poster,omitempty" validate:"required,min=2,max=32"`
	WatchlistID primitive.ObjectID `json:"watchlist,omitempty" bson:"watchlist,omitempty"`
	TmdbID      int                `json:"tmdbId,omitempty" bson:"tmdbId,omitempty"`
}
