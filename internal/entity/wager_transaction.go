package entity

import "time"

type WagerTransaction struct {
	ID                    string
	WalletID              string
	ProviderTransactionID string
	TransactionType       string
	Amount                int64
	Status                string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}