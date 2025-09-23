-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (gen_random_uuid(), now(), now(), @email, @password)
RETURNING *;

-- name: RemoveAllUsers :exec
Delete from users;

-- name: GetByEmail :one
Select email, hashed_password, id, created_at, updated_at from users
where email = @email;

-- name: UpdateUser :one
update users set email = @email, hashed_password = @password, updated_at = now()
where id = @id
RETURNING *;

-- name: UpgradeUser :one
update users set is_chirpy_red = @is_chirpy_red
where id = @id
RETURNING id;