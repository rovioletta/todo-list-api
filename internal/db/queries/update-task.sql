-- name: UpdateTask :one
UPDATE
  tasks
SET
  title = COALESCE(sqlc.narg('new_title'), title),
  description = COALESCE(sqlc.narg('new_description'), description),
  status = COALESCE(sqlc.narg('status') :: task_status, status)
WHERE
  id = sqlc.arg('id') :: integer
RETURNING id;