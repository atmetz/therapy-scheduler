package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	db *sql.DB
}

// Create new client
func NewClient(pathToDB string) (Client, error) {
	// set db to database path
	db, err := sql.Open("sqlite3", pathToDB)
	if err != nil {
		return Client{}, err
	}
	// create client
	c := Client{db}
	// Migrate database tables
	err = c.autoMigrate()
	if err != nil {
		return Client{}, err
	}
	return c, nil
}

// Set Database Schema
func (c *Client) autoMigrate() error {
	// create provider table
	providerTable := `
	CREATE TABLE IF NOT EXISTS providers (
		id UUID PRIMARY KEY,
		name TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		pasword TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL
	);
	`
	// create userTable in database
	_, err := c.db.Exec(providerTable)
	if err != nil {
		return err
	}

	// create client table
	platformTable := `
	CREATE TABLE IF NOT EXISTS clients (
		id UUID PRIMARY KEY,
		name TEXT NOT NULL,	
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	// create userTable in database
	_, err = c.db.Exec(platformTable)
	if err != nil {
		return err
	}

	// create client table
	clientTable := `
	CREATE TABLE IF NOT EXISTS clients (
		id UUID PRIMARY KEY,
		name TEXT NOT NULL,
		phone_number TEXT,
		email TEXT UNIQUE NOT NULL,		
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		provider_id UUID,
		platform_id UUID,
		FOREIGN KEY(provider_id) REFERENCES providers(id),
		FOREIGN KEY(platmform_id) REFERENCES platform(id)
	);
	`
	// create userTable in database
	_, err = c.db.Exec(clientTable)
	if err != nil {
		return err
	}
	return nil
}
