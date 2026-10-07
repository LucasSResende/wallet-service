package dto

type CreateWalletRequest struct {
	PlayerID string `json:"playerId"`
	Currency string `json:"currency"`
}