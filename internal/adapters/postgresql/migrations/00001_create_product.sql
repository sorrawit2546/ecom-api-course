-- +goose Up
CREATE TABLE IF NOT EXISTS products (
    id  BIGSERIAL PRIMARY KEY,
    name TEXT,
    price_in_cents  INTEGER  NOT NULL CHECK(price_in_cent >= 0),
    quantity    INTEGER NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);   


-- +goose Down
DROP TABLE IF EXISTS products;
