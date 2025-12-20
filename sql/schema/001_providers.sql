-- +goose Up
CREATE TABLE providers (
		id UUID PRIMARY KEY,
		name TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL,
		password TEXT NOT NULL,
		phone_number TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		sessions_available INTEGER NOT NULL
	);

-- +goose Down
DROP TABLE providers;