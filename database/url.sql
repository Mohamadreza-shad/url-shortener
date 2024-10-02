-- name: ShortenUrl :one
INSERT INTO urls (
    long_url,
    short_url,
    expired_at
) VALUES (
    $1, $2, $3
) returning *;

