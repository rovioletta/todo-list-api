-- name: GetUserByLogin :one
SELECT login, password_hash FROM users
WHERE login = @login :: string
LIMIT 1;