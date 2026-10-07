package entity

import "time"

type Wallet struct {
	ID        string    `json:"id"`
	PlayerID  string    `json:"playerId"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	Version   int64     `json:"version"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}