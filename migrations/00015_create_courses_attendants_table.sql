-- +goose NO TRANSACTION
-- +goose Up
CREATE TABLE IF NOT EXISTS human_resources.courses_attendants (
    course_id INT NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    PRIMARY KEY (course_id, user_id)
);

-- +goose Down
DROP TABLE IF EXISTS human_resources.courses_attendants;
