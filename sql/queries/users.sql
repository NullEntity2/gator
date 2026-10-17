-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUserByName :one
SELECT id, created_at, updated_at, name
FROM users
WHERE name = $1;

-- name: DeleteUsers :exec
TRUNCATE TABLE users CASCADE;

-- name: GetUsers :many
SELECT id, created_at, updated_at, name
FROM users;