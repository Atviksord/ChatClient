-- name: LoginUser :one
SELECT * from users
WHERE username = $1 AND password = $2;