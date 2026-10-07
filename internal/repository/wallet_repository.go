package repository

import "wallet-service/internal/entity"

type WalletRepository interface {

	Create(wallet *entity.Wallet) error

	FindByID(id string) (*entity.Wallet,error)

	Update(wallet *entity.Wallet) error
}