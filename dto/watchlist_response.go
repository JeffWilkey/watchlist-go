package dto

import (
	"github.com/jeffwilkey/watchlist-go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WatchlistResponse struct {
	ID            primitive.ObjectID   `json:"id"`
	Name          string               `json:"name"`
	Description   string               `json:"description"`
	OwnerID       primitive.ObjectID   `json:"ownerId"`
	Collaborators []primitive.ObjectID `json:"collaborators"`
}

type WatchlistWithMediaResponse struct {
	ID            primitive.ObjectID   `json:"id"`
	Name          string               `json:"name"`
	Description   string               `json:"description"`
	OwnerID       primitive.ObjectID   `json:"ownerId"`
	Collaborators []primitive.ObjectID `json:"collaborators"`
	Media         []MediaResponse      `json:"media"`
}

func CreateWatchlistResponse(watchlist model.Watchlist) WatchlistResponse {
	return WatchlistResponse{
		ID:            watchlist.ID,
		Name:          watchlist.Name,
		Description:   watchlist.Description,
		OwnerID:       watchlist.OwnerID,
		Collaborators: watchlist.Collaborators,
	}
}

func CreateWatchlistWithMediaResponse(watchlist model.Watchlist, media []MediaResponse) WatchlistWithMediaResponse {
	return WatchlistWithMediaResponse{
		ID:            watchlist.ID,
		Name:          watchlist.Name,
		Description:   watchlist.Description,
		OwnerID:       watchlist.OwnerID,
		Collaborators: watchlist.Collaborators,
		Media:         media,
	}
}

func CreateWatchlistsResponse(watchlists []model.Watchlist) []WatchlistResponse {
	watchlistsResponse := make([]WatchlistResponse, 0)
	for _, watchlist := range watchlists {
		watchlistsResponse = append(watchlistsResponse, CreateWatchlistResponse(watchlist))
	}
	return watchlistsResponse
}
