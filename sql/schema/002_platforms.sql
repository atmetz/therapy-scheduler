-- +goose Up
CREATE TABLE platforms (
		id UUID PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,	
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL
	);

-- +goose Down
DROP TABLE platforms;