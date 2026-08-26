-- name: GetUserByLogin :one
SELECT id, login, password_hash FROM users
WHERE login = @login :: text
LIMIT 1;