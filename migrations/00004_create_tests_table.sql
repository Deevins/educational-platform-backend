-- +goose NO TRANSACTION
-- +goose Up

CREATE TABLE IF NOT EXISTS human_resources.tests (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    section_id INT NOT NULL,
    serial_number INT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);


-- +goose Down

DROP TABLE if exists human_resources.tests;
