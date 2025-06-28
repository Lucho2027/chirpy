-- +goose Up
CREATE TABLE
	chirps(
		id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
		message text not null,
		user_id uuid not null references users(id) on delete cascade,
		created_at TIMESTAMP DEFAULT now(),
		updated_at TIMESTAMP DEFAULT now()
	);

-- +goose Down
DROP TABLE chirps;