-- Enable UUID support (PostgreSQL only)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Table: consumers
CREATE TABLE IF NOT EXISTS consumers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nik VARCHAR(100) UNIQUE NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    legal_name VARCHAR(255),
    place_of_birth VARCHAR(100),
    date_of_birth DATE,
    salary NUMERIC,
    photo_ktp TEXT,
    photo_selfie TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: limits
CREATE TABLE IF NOT EXISTS limits (
    id SERIAL PRIMARY KEY,
    consumer_id UUID NOT NULL,
    tenor INTEGER NOT NULL, -- in months (e.g. 1, 2, 3, 6)
    amount NUMERIC NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_consumer FOREIGN KEY(consumer_id) REFERENCES consumers(id) ON DELETE CASCADE,
    CONSTRAINT unique_consumer_tenor UNIQUE(consumer_id, tenor)
);

-- Table: transactions
CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    consumer_id UUID NOT NULL,
    contract_number VARCHAR(100) UNIQUE NOT NULL,
    tenor INTEGER NOT NULL,
    otr NUMERIC NOT NULL,
    admin_fee NUMERIC,
    installment NUMERIC,
    interest NUMERIC,
    asset_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_transaction_consumer FOREIGN KEY(consumer_id) REFERENCES consumers(id) ON DELETE CASCADE
);
