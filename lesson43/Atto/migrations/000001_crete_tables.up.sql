CREATE TABLE IF NOT EXISTS cards (
    card_id UUID PRIMARY KEY,
    card_number VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS stations (
    station_id UUID PRIMARY KEY,
    station_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS transactions (
    transaction_id UUID PRIMARY KEY,
    card_id UUID NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    terminal_id UUID DEFAULT NULL,
    transaction_type ENUM('credit', 'debit') NOT NULL,
    transaction_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (card_id) REFERENCES card(card_id),
    FOREIGN KEY (terminal_id) REFERENCES station(station_id)
);
