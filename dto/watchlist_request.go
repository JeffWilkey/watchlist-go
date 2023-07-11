package dto

type WatchlistUpdateRequest struct {
	Name        string `json:"name" bson:"name,omitempty" validate:"min=1,max=32"`
	Description string `json:"description" bson:"description,omitempty" validate:"max=256"`
}
