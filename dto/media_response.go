package dto

import (
	"github.com/jeffwilkey/watchlist-go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MediaResponse struct {
	ID          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Year        string             `json:"year"`
	Poster      string             `json:"poster"`
	WatchlistID primitive.ObjectID `json:"watchlist"`
	TmdbID      int                `json:"tmdbId"`
}

func CreateMediaResponse(media model.Media) MediaResponse {
	return MediaResponse{
		ID:          media.ID,
		Title:       media.Title,
		Year:        media.Year,
		Poster:      media.Poster,
		WatchlistID: media.WatchlistID,
		TmdbID:      media.TmdbID,
	}
}

func CreateMediaListResponse(media []model.Media) []MediaResponse {
	mediaResponse := make([]MediaResponse, 0)
	for _, media := range media {
		mediaResponse = append(mediaResponse, CreateMediaResponse(media))
	}
	return mediaResponse
}
