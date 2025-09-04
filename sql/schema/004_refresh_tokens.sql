-- +goose Up
CREATE TABLE
	refresh_tokens (
		token text PRIMARY KEY NOT Null,
		created_at TIMESTAMP DEFAULT now(),
		updated_at TIMESTAMP DEFAULT now(),
		user_id uuid not null references users(id) on delete cascade,
		expires_at TIMESTAMP DEFAULT NULL,
		revoked_at TIMESTAMP DEFAULT NULL
	);

-- +goose Down
DROP TABLE refresh_tokens;