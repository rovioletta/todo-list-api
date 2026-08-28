-- name: GetTaskByID :one
SELECT
  *
FROM
  tasks
WHERE
  id = @id :: integer
LIMIT
  1;