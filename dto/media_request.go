package dto

type MediaUpdateRequest struct {
	Status string `json:"status" bson:"status,omitempty" validate:"oneof=watched unwatched in-progress"`
}
