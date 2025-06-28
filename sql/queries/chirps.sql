-- name: CreateChirp :one
INSERT INTO chirps (id, message, user_id, created_at, updated_at)
VALUES (gen_random_uuid(), @message, @user_id, now(), now())
RETURNING *;

-- name: RemoveAllChirps :exec
TRUNCATE TABLE chirps RESTART IDENTITY; 

-- name: GetAllChirps :many
SELECT * FROM chirps
ORDER By created_at ASC;