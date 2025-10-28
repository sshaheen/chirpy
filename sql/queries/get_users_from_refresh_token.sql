-- name: GetUsersFromRefreshToken :one
SELECT *
FROM users
LEFT JOIN refresh_tokens
ON id = user_id;