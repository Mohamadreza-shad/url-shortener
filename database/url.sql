-- name: ShortenUrl :one
INSERT INTO urls (
    long_url,
    short_url,
    expired_at
) VALUES (
    $1, $2, $3
) returning *;

-- name: UrlByLongUrl :one
SELECT * FROM urls WHERE long_url = $1;

-- name: UrlByShortUrl :one
SELECT * FROM urls WHERE short_url = $1;