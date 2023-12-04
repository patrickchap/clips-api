
-- name: CreateLike :one
INSERT INTO likes (
  video_id,
  user_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetLikesByVideo :many
SELECT * 
FROM 
  likes 
WHERE
  video_id = sqlc.arg(video_id)
ORDER BY id 
LIMIT $1
OFFSET $2;

-- name: DeleteLikesByVideoAndUser :exec
Delete
FROM 
  likes 
WHERE
  video_id = sqlc.arg(video_id)
AND
  user_id = sqlc.arg(user_id);
