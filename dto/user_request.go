package dto

type UserUpdateRequest struct {
	FirstName string `json:"firstName" bson:"firstName,omitempty" validate:"required,min=1,max=32"`
	LastName  string `json:"lastName" bson:"lastName,omitempty" validate:"required,min=2,max=32"`
}
