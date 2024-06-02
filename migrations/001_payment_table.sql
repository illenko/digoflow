-- +goose Up
CREATE TABLE IF NOT EXISTS payment
(
    id             UUID PRIMARY KEY,
    description    VARCHAR(256),
    amount         NUMERIC(10, 2),
    card           VARCHAR(256),
    payer_details  jsonb,
    payment_system jsonb
);