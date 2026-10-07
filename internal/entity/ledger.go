package entity

import "time"

type Ledger struct {
	ID              string    `json:"id"`
	WalletID        string    `json:"walletId"`
	TransactionID   string    `json:"transactionId"`
	TransactionType string    `json:"transactionType"`
	Amount          int64     `json:"amount"`
	BalanceBefore   int64     `json:"balanceBefore"`
	BalanceAfter    int64     `json:"balanceAfter"`
	CreatedAt       time.Time `json:"createdAt"`
}