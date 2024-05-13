-- +goose NO TRANSACTION
-- +goose Up
CREATE TABLE IF NOT EXISTS human_resources.courses_lectures (
    id SERIAL PRIMARY KEY,
    course_id INT NOT NULL,
    lecture_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS human_resources.courses_lectures;
