-- name: CreateChirp :one
INSERT INTO chirps (id, message, user_id, created_at, updated_at)
VALUES (gen_random_uuid(), @message, @user_id, now(), now())
RETURNING *;

-- name: RemoveAllChirps :exec
TRUNCATE TABLE chirps RESTART IDENTITY; 

-- name: GetAllChirps :many
SELECT * FROM chirps
ORDER By created_at ASC;

-- name: GetChirpById :one
select * from chirps where id = @id;

-- name: DeleteChirpById :exec
Delete from chirps where id = @chirp_id and user_id = @user_id;

-- name: GetAllChirpsByAuthor :many
SELECT * FROM chirps
WHERE user_id = @user_id
ORDER BY created_at ASC;