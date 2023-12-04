-- name: CreateComment :one
INSERT INTO comments (
  content,
  video_id,
  user_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetCommentsByVideo :many
SELECT * 
FROM 
  comments 
WHERE
  video_id = sqlc.arg(video_id)
ORDER BY id 
LIMIT $1
OFFSET $2;

-- name: DeleteCommentsByVideo :exec
Delete
FROM 
  comments 
WHERE
  video_id = sqlc.arg(video_id);
