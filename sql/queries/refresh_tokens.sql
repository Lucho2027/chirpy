-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expires_at, revoked_at)
VALUES (@token, now(), now(), @user_id, @expires_at, null)
RETURNING *;