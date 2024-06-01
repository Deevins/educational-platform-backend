-- +goose NO TRANSACTION
-- +goose Up

INSERT INTO human_resources.skill_levels (name, created_at, updated_at)
    VALUES ('Начальный', now(), now()),
           ('Средний', now(), now()),
           ('Продвинутый', now(), now());
-- +goose Down
