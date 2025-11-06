package models

import "time"

type User struct {
	ID		int       `json:"ID"`
	Name	string    `json:"name"  binding:"requiredn min=10, max=200"`
	Email	string    `json:"email" binding:"required.email"`
	CreatedAt	time.Time `json:"createdAt"`
	UpdatedAt	time.Time `json:"updatedAt"`
}