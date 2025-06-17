-- +goose Up
CREATE TABLE
	users (
		id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
		created_at TIMESTAMP DEFAULT now(),
		updated_at TIMESTAMP DEFAULT now(),
		email text not null unique
	);

-- +goose Down
DROP TABLE users;