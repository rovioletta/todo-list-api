-- name: GetUserByLogin :one
SELECT login, password_hash FROM users
WHERE login = @login :: text
LIMIT 1;