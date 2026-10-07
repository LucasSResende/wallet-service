package service

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/lucassresende/wallet-service/internal/entity"
	"github.com/lucassresende/wallet-service/internal/repository"
)

type WalletService struct {
	repository repository.WalletRepository
}

func NewWalletService(
	repository repository.WalletRepository,
) *WalletService {
	return &WalletService{
		repository: repository,
	}
}

func (s *WalletService) CreateWallet(
	playerID string,
	currency string,
) (*entity.Wallet, error) {

	now := time.Now()

	wallet := &entity.Wallet{
		ID:        uuid.New().String(),
		PlayerID:  playerID,
		Balance:   0,
		Currency:  currency,
		Version:   1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := s.repository.Create(wallet)

	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (s *WalletService) GetWallet(
	id string,
) (*entity.Wallet, error) {

	return s.repository.GetByID(id)
}

func (s *WalletService) Bet(
	id string,
	amount int64,
) error {

	wallet, err := s.repository.GetByID(id)

	if err != nil {
		return err
	}

	if wallet.Balance < amount {
		return errors.New("insufficient balance")
	}

	wallet.Balance -= amount
	wallet.Version++
	wallet.UpdatedAt = time.Now()

	return s.repository.Update(wallet)
}

func (s *WalletService) Win(
	id string,
	amount int64,
) error {

	wallet, err := s.repository.GetByID(id)

	if err != nil {
		return err
	}

	wallet.Balance += amount
	wallet.Version++
	wallet.UpdatedAt = time.Now()

	return s.repository.Update(wallet)
}