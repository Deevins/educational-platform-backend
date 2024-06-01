-- +goose NO TRANSACTION
-- +goose Up

CREATE TABLE IF NOT EXISTS human_resources.tests (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);


-- +goose Down

DROP TABLE if exists human_resources.tests;
