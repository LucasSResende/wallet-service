CREATE TABLE IF NOT EXISTS wallet_ledger
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    wallet_id UUID NOT NULL,

    transaction_id UUID NOT NULL,

    transaction_type VARCHAR(50) NOT NULL,

    amount BIGINT NOT NULL,

    balance_before BIGINT NOT NULL,

    balance_after BIGINT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_wallet_ledger_wallet
        FOREIGN KEY(wallet_id)
            REFERENCES wallets(id)
);
            
CREATE INDEX idx_wallet_ledger_wallet_id
ON wallet_ledger(wallet_id);

CREATE INDEX idx_wallet_ledger_transaction_id
ON wallet_ledger(transaction_id);