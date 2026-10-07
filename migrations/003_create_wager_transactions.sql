CREATE TABLE IF NOT EXISTS wager_transactions
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    wallet_id UUID NOT NULL,

    provider_transaction_id UUID NOT NULL,

    transaction_type VARCHAR(20) NOT NULL,

    amount BIGINT NOT NULL,

    status VARCHAR(20) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_wager_wallet
        FOREIGN KEY(wallet_id)
            REFERENCES wallets(id)
);

CREATE UNIQUE INDEX idx_provider_transaction
ON wager_transactions(provider_transaction_id);