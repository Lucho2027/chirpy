-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expires_at, revoked_at)
VALUES (@token, now(), now(), @user_id, @expires_at, null)
RETURNING *;

-- name: GetUserFromRefreshToken :one
Select user_id from refresh_tokens
where token = @token
and revoked_at IS NULL and expires_at > now();

-- name: RevokeToken :exec
update refresh_tokens set updated_at = now(), revoked_at = now()
where token = @token;