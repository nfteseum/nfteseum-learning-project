-- name: GetUser :one
SELECT * FROM users WHERE addr = $1;

-- name: CreateUser :one
INSERT INTO users (addr, name, random_msg) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateUser :one
UPDATE users SET name = $2, pfp=$3, random_msg=$4 WHERE addr = $1 RETURNING *;
