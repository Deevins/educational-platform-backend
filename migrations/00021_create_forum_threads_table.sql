-- +goose NO TRANSACTION
-- +goose Up

CREATE TABLE human_resources.forum_threads (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);


-- +goose Down

DROP TABLE IF EXISTS human_resources.forum_threads;
