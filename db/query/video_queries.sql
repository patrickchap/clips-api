-- name: CreateVideo :one
INSERT INTO videos (
  title,
  description,
  file_url,
  thumbnail_url,
  user_id
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetVideo :one
SELECT * FROM videos 
WHERE id = $1 LIMIT 1;


-- name: GetVideoWithLikesWithSearch :one
SELECT
    v.*,
    COUNT(l.id) AS like_count
FROM
    videos v
LEFT JOIN
    likes l ON v.id = l.video_id
WHERE
    v.id = $1
GROUP BY
    v.id;

-- name: GetUserVideoWithLikes :many
SELECT
    v.*,
    COUNT(l.id) AS like_count
FROM
    videos v
LEFT JOIN
    likes l ON v.id = l.video_id
WHERE
    v.user_id = sqlc.arg(user_id)::text
GROUP BY
    v.id
ORDER BY 
    v.created_at desc
LIMIT $1
OFFSET $2;

-- name: GetVideoWithLikes :one
SELECT
    v.*,
    COUNT(l.id) AS like_count
FROM
    videos v
LEFT JOIN
    likes l ON v.id = l.video_id
WHERE
    v.id = $1
GROUP BY
    v.id;

-- name: ListVideos :many
SELECT * FROM videos 
ORDER BY id 
LIMIT $1
OFFSET $2;

-- name: ListVideosWithLikesAndSearch :many
SELECT
    v.*,
    COUNT(l.id) AS like_count
FROM
    videos v
LEFT JOIN
    likes l ON v.id = l.video_id
WHERE
    v.title ILIKE sqlc.arg(search)::text
OR
    v.description ILIKE sqlc.arg(search)::text
GROUP BY
    v.id
ORDER BY 
	v.created_at desc

LIMIT $1
OFFSET $2;

-- name: DeleteVideo :exec
DELETE FROM videos 
WHERE id = $1;
