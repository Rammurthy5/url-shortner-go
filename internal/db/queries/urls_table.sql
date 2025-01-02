-- name: InsertUrl :one
INSERT INTO url_mapping (long_url, short_url)
VALUES ($1, $2) RETURNING *;

-- name: GetUrl :one
SELECT from url_mapping
WHERE long_url = $1 LIMIT 1;

-- name: UpdateUrl :one
UPDATE url_mapping SET short_url = $1
WHERE long_url = $1 RETURNING *;

-- name: DeleteUrl :exec
DELETE from url_mapping
WHERE long_url = $1;
