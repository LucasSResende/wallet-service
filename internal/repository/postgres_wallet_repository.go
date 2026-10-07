package repository

import (
	"database/sql"

	"github.com/lucassresende/wallet-service/internal/entity"
)

type PostgresWalletRepository struct {
	DB *sql.DB
}

func NewPostgresWalletRepository(db *sql.DB) *PostgresWalletRepository {
	return &PostgresWalletRepository{DB: db}
}

func (r *PostgresWalletRepository) Create(wallet *entity.Wallet) error {

	query := `
	INSERT INTO wallets
	(
		id,
		player_id,
		balance,
		currency,
		version,
		created_at,
		updated_at
	)
	VALUES
	($1,$2,$3,$4,$5,$6,$7)
	`

	_, err := r.DB.Exec(
		query,
		wallet.ID,
		wallet.PlayerID,
		wallet.Balance,
		wallet.Currency,
		wallet.Version,
		wallet.CreatedAt,
		wallet.UpdatedAt,
	)

	return err
}

func (r *PostgresWalletRepository) GetByID(id string) (*entity.Wallet, error) {

	query := `
	SELECT
		id,
		player_id,
		balance,
		currency,
		version,
		created_at,
		updated_at
	FROM wallets
	WHERE id = $1
	`

	var wallet entity.Wallet

	err := r.DB.QueryRow(query, id).Scan(
		&wallet.ID,
		&wallet.PlayerID,
		&wallet.Balance,
		&wallet.Currency,
		&wallet.Version,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (r *PostgresWalletRepository) Update(wallet *entity.Wallet) error {

	query := `
	UPDATE wallets
	SET
		balance = $1,
		version = $2,
		updated_at = $3
	WHERE id = $4
	`

	_, err := r.DB.Exec(
		query,
		wallet.Balance,
		wallet.Version,
		wallet.UpdatedAt,
		wallet.ID,
	)

	return err
}