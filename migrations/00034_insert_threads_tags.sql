-- +goose NO TRANSACTION
-- +goose Up

INSERT INTO human_resources.threads_tags (name)
VALUES  ('ReactJS'),
        ('TypeScript'),
        ('JavaScript'),
        ('NodeJS'),
        ('Python'),
        ('Java'),
        ('Golang'),
        ('Docker'),
        ('Kubernetes'),
        ('Ansible'),
        ('Kafka'),
        ('Jenkins'),
        ('GitLab CI/CD'),
        ('Unit Testing'),
        ('Интеграционное тестирование'),
        ('E2E тестированеи'),
        ('Тестирование безопасности'),
        ('Тестирование проникнований'),
        ('Аудит безопасности');

-- +goose Down
