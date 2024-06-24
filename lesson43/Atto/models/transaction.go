package models

import "time"

type TransactionResponse struct {
	Transaction_ID   string    `json:"transaction_id"`
	Card_ID          string    `json:"card_id"`
	Amount           float64   `json:"amount"`
	Terminal_ID      string    `json:"terminal_id"`
	Transaction_Type string    `json:"transaction_type"`
	Transaction_Time time.Time `json:"transaction_time"`
}

type TransactionRequest struct {
	Transaction_ID   string  `json:"transaction_id"`
	Card_ID          string  `json:"card_id"`
	Amount           float64 `json:"amount"`
	Terminal_ID      string  `json:"terminal_id"`
	Transaction_Type string  `json:"transaction_type"`
}
