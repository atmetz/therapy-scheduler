-- name: CreateClient :one
INSERT INTO clients (id, created_at, updated_at, name, phone_number, email, platform_id, provider_id)
VALUES (
    ?, 
    ?, 
    ?, 
    ?, 
    ?, 
    ?,
    ?,
    ?
)
RETURNING *;