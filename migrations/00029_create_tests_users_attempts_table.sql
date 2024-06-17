-- +goose NO TRANSACTION
-- +goose Up

CREATE TABLE IF NOT EXISTS human_resources.tests_users_attempts (
    id SERIAL PRIMARY KEY,
    test_id INT NOT NULL,
    user_id INT NOT NULL,
    attempt_number INT NOT NULL,
    total_questions_count INT NOT NULL,
    correct_answers_count INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);


-- +goose Down

DROP TABLE if exists human_resources.tests_users_attempts;
