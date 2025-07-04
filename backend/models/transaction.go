package models

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	Category  string    `json:"category"`
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	Type      string    `json:"type"` 
	CreatedAt time.Time `json:"created_at"`
}
