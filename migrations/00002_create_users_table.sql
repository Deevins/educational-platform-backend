-- +goose NO TRANSACTION
-- +goose Up

CREATE TABLE IF NOT EXISTS human_resources.users (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    description TEXT,
    avatar_url TEXT default '',
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hashed VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    has_user_tried_instructor BOOLEAN DEFAULT FALSE,
    phone_number VARCHAR(255) NOT NULL,
    role human_resources.roles DEFAULT 'USER',
    students_count INTEGER DEFAULT 0 not null,
    courses_count INTEGER DEFAULT 0 not null,
    instructor_rating INTEGER DEFAULT 0 not null
);


-- +goose Down
DROP TABLE human_resources.users;
