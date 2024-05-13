-- +goose NO TRANSACTION
-- +goose Up

CREATE TABLE human_resources.threads_tags_links (
    id SERIAL PRIMARY KEY,
    thread_id INT NOT NULL,
    tag_id INT NOT NULL
);


-- +goose Down

DROP TABLE IF EXISTS human_resources.threads_tags_links;
