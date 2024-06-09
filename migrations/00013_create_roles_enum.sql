-- +goose NO TRANSACTION
-- +goose Up
CREATE TYPE human_resources.roles AS ENUM ('ADMIN','MODERATOR', 'USER');

-- +goose Down
DROP TYPE IF EXISTS human_resources.roles;
