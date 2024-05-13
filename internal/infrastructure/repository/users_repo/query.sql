-- name: GetUsers :many
SELECT * from human_resources.users;


-- name: GetUserByID :one
SELECT * from human_resources.users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO human_resources.users (full_name, email,description, password_hashed, phone_number) VALUES ($1, $2, $3, $4, $5) RETURNING id;

-- name: GetUserByEmailAndHashedPassword :one
SELECT * from human_resources.users WHERE email = $1 AND password_hashed = $2;