-- +goose NO TRANSACTION
-- +goose Up

CREATE TABLE human_resources.threads_tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);


-- +goose Down

DROP TABLE IF EXISTS human_resources.threads_tags;
