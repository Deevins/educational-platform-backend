-- +goose NO TRANSACTION
-- +goose Up

INSERT INTO human_resources.languages (name, created_at, updated_at)
    values ('Русский', now(), now()),
           ('Английский', now(),now());


-- +goose Down
