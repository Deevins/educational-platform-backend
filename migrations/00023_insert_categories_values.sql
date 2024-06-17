-- +goose NO TRANSACTION
-- +goose Up

INSERT INTO human_resources.categories (name, created_at, updated_at)
        VALUES ('Разработка интерфейсов (frontend)', now(), now()),
               ('Разработка сервисов (backend)', now(), now()),
               ('Разработка комплексных приложений (fullstack)', now(), now()),
               ('Инфраструктура', now(), now()),
               ('Тестирование', now(), now()),
               ('Безопасность', now(), now()),
               ('Другое', now(), now());


-- +goose Down
