-- +goose NO TRANSACTION
-- +goose Up
CREATE TABLE IF NOT EXISTS human_resources.courses_lectures (
    id SERIAL PRIMARY KEY,
    section_id INT NOT NULL,
    lecture_id INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    PRIMARY KEY (section_id, lecture_id)
);

-- +goose Down
DROP TABLE IF EXISTS human_resources.courses_lectures;
