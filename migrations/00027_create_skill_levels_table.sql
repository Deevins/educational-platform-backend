-- +goose NO TRANSACTION
-- +goose Up

CREATE TABLE human_resources.skill_levels (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- +goose Down

DROP TABLE IF EXISTS human_resources.skill_levels;
