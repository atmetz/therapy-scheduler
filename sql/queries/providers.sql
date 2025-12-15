-- name: CreateProvider :one
INSERT INTO providers (id, created_at, updated_at, name, password, phone_number, email)
VALUES (
    ?, 
    ?, 
    ?, 
    ?, 
    ?, 
    ?,
    ?
)
RETURNING *;

-- name: GetProvider :one
SELECT * FROM providers
WHERE name = ?;

-- name: GetProviderByEmail :one
SELECT * FROM providers
WHERE email = ?;

-- name: GetProviders :many
SELECT * FROM providers;

-- name: GetProviderById :one
SELECT * FROM providers 
WHERE id = ?;