-- name: GetMessage :one
SELECT * FROM messages
WHERE message_id = $1 LIMIT 1;

-- name: ListMessages :many
SELECT * FROM messages
ORDER BY created_at desc;

-- name: CreateMessage :one
INSERT INTO messages (
  user_id, message_text, message_type
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteMessage :exec
DELETE FROM messages
WHERE message_id = $1;