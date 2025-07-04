package models

import "time"

type Goal struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	TargetAmount  float64   `json:"target_amount"`
	CurrentAmount float64   `json:"current_amount"`
	CreatedAt     time.Time `json:"created_at"`
}
