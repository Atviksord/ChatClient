-- name: AddApikey :one
UPDATE users SET api_key = $1 WHERE username = $2
RETURNING id, username, api_key;
