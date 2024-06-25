package models

import "time"

type CardResponse struct {
	Card_ID     string    `json:"card_id"`
	Card_Number string    `json:"card_number"`
	User_id     string    `json:"user_id"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}