package dto

type WalletResponse struct {
	ID       string `json:"id"`
	PlayerID string `json:"playerId"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}