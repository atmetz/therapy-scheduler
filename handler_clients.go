package main

import (
	"context"
	"fmt"
	"time"

	"github.com/atmetz/therapy-scheduler/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerClientsCreate(params parameters) (database.Client, error) {

	user, err := cfg.db.CreateClient(context.Background(), database.CreateClientParams{
		ID:          uuid.New(),
		Name:        params.Name,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Email:       params.Email,
		PhoneNumber: params.Phone,
		PlatformID:  params.PlatformID,
		ProviderID:  params.ProviderID,
	})
	if err != nil {
		return database.Client{}, fmt.Errorf("couldn't create client %s", err)
	}

	return user, nil

}
