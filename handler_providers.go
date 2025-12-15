package main

import (
	"context"
	"fmt"
	"time"

	"github.com/atmetz/therapy-scheduler/internal/auth"
	"github.com/atmetz/therapy-scheduler/internal/database"
	"github.com/google/uuid"
)

type parameters struct {
	Name     string
	Email    string
	Phone    string
	Password string
}

func (cfg *apiConfig) handlerProvidersCreate(params parameters) (database.Provider, error) {

	_, err := cfg.db.GetProviderByEmail(context.Background(), params.Email)
	if err == nil {
		return database.Provider{}, fmt.Errorf("provider already exists: %v", err)
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		return database.Provider{}, fmt.Errorf("couldn't hash password %s", err)
	}

	user, err := cfg.db.CreateProvider(context.Background(), database.CreateProviderParams{
		ID:          uuid.New(),
		Name:        params.Name,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Password:    hashedPassword,
		PhoneNumber: params.Phone,
		Email:       params.Email,
	})
	if err != nil {
		return database.Provider{}, fmt.Errorf("couldn't create provider %s", err)
	}

	return user, nil

}

func (cfg *apiConfig) handlerProviderLogin(params parameters) (database.Provider, error) {

	user, err := cfg.db.GetProviderByEmail(context.Background(), params.Email)
	if err != nil {
		return database.Provider{}, fmt.Errorf("cannot find provider: %v", err)
	}

	match, err := auth.CheckPasswordHash(params.Password, user.Password)

	if err != nil || !match {
		return database.Provider{}, fmt.Errorf("incorrect password %s", err)
	}

	return user, nil
}
