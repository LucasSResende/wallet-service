CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS wallets
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    player_id UUID NOT NULL,

    balance BIGINT NOT NULL DEFAULT 0,

    currency VARCHAR(3) NOT NULL DEFAULT 'BRL',

    version BIGINT NOT NULL DEFAULT 1,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_wallets_player_id
ON wallets(player_id);