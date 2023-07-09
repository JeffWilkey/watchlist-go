package dto

import (
	"github.com/jeffwilkey/watchlist-go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WatchListResponse struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	UserID      primitive.ObjectID `json:"userId"`
}

func CreateWatchListResponse(watchlist model.Watchlist) WatchListResponse {
	return WatchListResponse{
		ID:          watchlist.ID,
		Name:        watchlist.Name,
		Description: watchlist.Description,
		UserID:      watchlist.UserID,
	}
}

func CreateWatchListsResponse(watchlists []model.Watchlist) []WatchListResponse {
	watchlistsResponse := make([]WatchListResponse, 0)
	for _, watchlist := range watchlists {
		watchlistsResponse = append(watchlistsResponse, CreateWatchListResponse(watchlist))
	}
	return watchlistsResponse
}
