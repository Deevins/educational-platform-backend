-- +goose NO TRANSACTION
-- +goose Up

CREATE TABLE human_resources.forum_threads_messages (
    id SERIAL PRIMARY KEY,
    thread_id INT NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    user_sent_id INT NOT NULL
);


-- +goose Down

DROP TABLE IF EXISTS human_resources.forum_threads_messages;
