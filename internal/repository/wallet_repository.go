package repository

import "github.com/lucassresende/wallet-service/internal/entity"


type WalletRepository interface {
	Create(wallet *entity.Wallet) error
	GetByID(id string) (*entity.Wallet, error)
	Update(wallet *entity.Wallet) error
}