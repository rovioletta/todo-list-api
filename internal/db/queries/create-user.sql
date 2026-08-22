-- name: CreateUser :exec
INSERT INTO users (login, password_hash)
VALUES ($1, $2);