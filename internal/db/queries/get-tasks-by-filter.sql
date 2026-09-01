-- name: GetTasksByFilter :many
SELECT
  *
FROM
  tasks
WHERE
  -- by title
  (sqlc.narg('filter_search_title')::text IS NULL OR title LIKE CONCAT('%', sqlc.narg('filter_search_title')::text, '%'))

  -- by status
  AND (sqlc.narg('filter_search_status')::text IS NULL OR LOWER(status::text) = LOWER(sqlc.narg('filter_search_status')::text))

  -- by created_at
  AND (sqlc.narg('filter_created_from')::timestamp IS NULL OR created_at >= @filter_created_from)
  AND (sqlc.narg('filter_created_to')::timestamp IS NULL OR created_at <= @filter_created_to)
  
ORDER BY
  -- by title
  CASE WHEN sqlc.narg('order_by_field')::text = 'title' AND sqlc.narg('order_type')::text = 'asc' THEN title END ASC,
  CASE WHEN sqlc.narg('order_by_field')::text = 'title' AND sqlc.narg('order_type')::text = 'desc' THEN title END DESC,

  -- by status
  CASE WHEN sqlc.narg('order_by_field')::text = 'status' AND sqlc.narg('order_type')::text = 'asc' THEN status END ASC,
  CASE WHEN sqlc.narg('order_by_field')::text = 'status' AND sqlc.narg('order_type')::text = 'desc' THEN status END DESC,
  
  -- by created_at
  CASE WHEN sqlc.narg('order_by_field')::text = 'created_at' AND sqlc.narg('order_type')::text = 'asc' THEN created_at END ASC,
  CASE WHEN sqlc.narg('order_by_field')::text = 'created_at' AND sqlc.narg('order_type')::text = 'desc' THEN created_at END DESC,
  
  -- default
  created_at DESC
LIMIT
  sqlc.arg('limit')::integer OFFSET sqlc.arg('offset')::integer;