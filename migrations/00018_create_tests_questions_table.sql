-- +goose NO TRANSACTION
-- +goose Up
CREATE TABLE IF NOT EXISTS human_resources.tests_questions (
    id SERIAL PRIMARY KEY,
    test_id INT NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS human_resources.tests_questions;
