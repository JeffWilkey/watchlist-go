package dto

type WatchlistUpdateRequest struct {
	Name        string `json:"name" validate:"max=32"`
	Description string `json:"description" validate:"max=256"`
}
