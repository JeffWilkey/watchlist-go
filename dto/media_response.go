package dto

import (
	"github.com/jeffwilkey/watchlist-go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MediaResponse struct {
	ID          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Poster      string             `json:"poster"`
	ReleaseDate primitive.DateTime `json:"releaseDate"`
	Status      string             `json:"status"`
	TmdbID      int                `json:"tmdbId"`
	WatchlistID primitive.ObjectID `json:"watchlistId"`
	CreatedAt   primitive.DateTime `json:"createdAt"`
}

func CreateMediaResponse(media model.Media) MediaResponse {
	return MediaResponse{
		ID:          media.ID,
		Title:       media.Title,
		Poster:      media.Poster,
		ReleaseDate: media.ReleaseDate,
		Status:      media.Status,
		TmdbID:      media.TmdbID,
		WatchlistID: media.WatchlistID,
		CreatedAt:   media.CreatedAt,
	}
}

func CreateMediaListResponse(media []model.Media) []MediaResponse {
	mediaResponse := make([]MediaResponse, 0)
	for _, media := range media {
		mediaResponse = append(mediaResponse, CreateMediaResponse(media))
	}
	return mediaResponse
}
