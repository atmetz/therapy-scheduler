-- +goose Up
CREATE TABLE clients (
		id UUID PRIMARY KEY,
		name TEXT NOT NULL,
		phone_number TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,		
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL,
		provider_id UUID,
		platform_id UUID,
		FOREIGN KEY(provider_id) REFERENCES providers(id),
		FOREIGN KEY(platform_id) REFERENCES platforms(id)
	);

-- +goose Down
DROP TABLE clients;