-- +goose NO TRANSACTION
-- +goose Up

CREATE TABLE IF NOT EXISTS human_resources.users (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    description TEXT,
    avatar_url VARCHAR(255) default '',
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hashed VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    has_user_tried_instructor BOOLEAN DEFAULT FALSE,
    phone_number VARCHAR(255) NOT NULL
);


-- +goose Down
DROP TABLE human_resources.users;
