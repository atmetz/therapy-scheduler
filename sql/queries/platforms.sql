-- name: CreatePlatform :one
INSERT INTO platforms (id, created_at, updated_at, name)
VALUES (
    ?, 
    ?, 
    ?, 
    ?
)
RETURNING *;

-- name: GetPlatforms :many
SELECT * FROM platforms;

-- name: GetPlatformByName :one
SELECT * FROM platforms
WHERE name = ?;