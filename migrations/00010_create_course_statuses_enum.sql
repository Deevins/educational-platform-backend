-- +goose NO TRANSACTION
-- +goose Up
CREATE TYPE human_resources.course_statuses AS ENUM ('DRAFT','PENDING', 'READY');

-- +goose Down
DROP TYPE IF EXISTS human_resources.course_statuses;
