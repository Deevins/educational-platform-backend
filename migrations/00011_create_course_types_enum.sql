-- +goose NO TRANSACTION
-- +goose Up
CREATE TYPE human_resources.course_types AS ENUM ('course','practice_course');

-- +goose Down
DROP TYPE IF EXISTS human_resources.course_types;
