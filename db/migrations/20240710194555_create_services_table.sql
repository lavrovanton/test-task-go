-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS services (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) UNIQUE NOT NULL,
    type VARCHAR(255) NOT NULL,
    payment_type VARCHAR(255) NOT NULL,
    price INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS services;
-- +goose StatementEnd
