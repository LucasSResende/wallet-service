package entity

import "time"

type Wallet struct {
	ID string

	PlayerID string

	Balance int64

	Currency string

	Version int64

	CreatedAt time.Time

	UpdatedAt time.Time
}