-- +goose NO TRANSACTION
-- +goose Up

INSERT INTO human_resources.subcategories (name, category_id, created_at, updated_at)
    VALUES  ('ReactJS', 1, now(), now()),
            ('TypeScript', 1, now(), now()),
            ('JavaScript', 1, now(), now()),
            ('NodeJS', 2, now(), now()),
            ('Python', 2, now(), now()),
            ('Java', 2, now(), now()),
            ('Golang', 2, now(), now()),
            ('ReactJS + NodeJS', 3, now(), now()),
            ('ReactJS + Python', 3, now(), now()),
            ('ReactJS + Java', 3, now(), now()),
            ('ReactJS + Golang', 3, now(), now()),
            ('VueJS + Golang', 3, now(), now()),
            ('Docker', 4, now(), now()),
            ('Kubernetes', 4, now(), now()),
            ('Ansible', 4, now(), now()),
            ('Kafka', 4, now(), now()),
            ('Jenkins', 4, now(), now()),
            ('GitLab CI/CD', 4, now(), now()),
            ('Unit Testing', 5, now(), now()),
            ('Интеграционное тестирование', 5, now(), now()),
            ('E2E тестированеи', 5, now(), now()),
            ('Тестирование безопасности', 6, now(), now()),
            ('Тестирование проникнований', 6, now(), now()),
            ('Аудит безопасности', 6, now(), now()),
            ('Другое', 7, now(), now());

-- +goose Down
