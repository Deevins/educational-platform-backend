-- +goose NO TRANSACTION
-- +goose Up
CREATE TABLE IF NOT EXISTS human_resources.lectures_resources (
    id SERIAL PRIMARY KEY,
    lecture_id INT NOT NULL,
    title text not null default '',
    extension text not null default '',
    resource_url text not null default '',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS human_resources.lectures_resources;
