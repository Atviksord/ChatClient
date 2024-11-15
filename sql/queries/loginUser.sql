-- name: SelectUser :one
SELECT * from users
WHERE username = $1;