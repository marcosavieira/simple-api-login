-- name: CreateUser :one
INSERT INTO users (
    username,
    password,
    email
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUsersByUsername :many
SELECT username FROM users WHERE LOWER(username) LIKE CONCAT('%', LOWER(@username::text), '%');

-- name: UpdateUser :one
UPDATE users SET
    username = COALESCE($1, username),
    password = COALESCE($2, password),
    email = COALESCE($3, email)
WHERE id = $4
RETURNING *;

-- name: DeleteUserID :exec
DELETE FROM users WHERE id = $1;

-- name: DeleteUserUsername :exec
DELETE FROM users WHERE username = $1;