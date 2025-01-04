-- name: InsertUrl :one
INSERT INTO url_mapping (long_url, short_url)
VALUES ($1, $2) RETURNING *;

-- name: GetUrl :one
SELECT * FROM url_mapping
WHERE long_url = $1 LIMIT 1;

-- name: UpdateUrl :one
UPDATE url_mapping SET short_url = $1
WHERE long_url = $2 RETURNING *;

-- name: DeleteUrl :exec
DELETE FROM url_mapping
WHERE long_url = $1;
