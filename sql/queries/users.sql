-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (gen_random_uuid(), now(), now(), @email, @password)
RETURNING *;

-- name: RemoveAllUsers :exec
Delete from users;

-- name: GetByEmail :one
Select email, hashed_password, id, created_at, updated_at from users
where email = @email;