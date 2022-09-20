package model

type Person struct {
	ID        uint   `json:"id,omitempty"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}
