package dto

type UserUpdateRequest struct {
	FirstName string `json:"firstName" validate:"required,min=1,max=32"`
	LastName  string `json:"lastName" validate:"required,min=2,max=32"`
}
