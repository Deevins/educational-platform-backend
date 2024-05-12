-- +goose NO TRANSACTION
-- +goose Up

CREATE SCHEMA IF NOT EXISTS human_resources;


-- +goose Down

DROP SCHEMA IF EXISTS human_resources;
