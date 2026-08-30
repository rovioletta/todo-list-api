-- name: DeleteTask :one
DELETE FROM tasks WHERE id = @id::integer
RETURNING id;