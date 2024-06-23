package models

import "time"

type UserResponse struct {
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRequest struct {
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Age       int       `json:"age"`
}
