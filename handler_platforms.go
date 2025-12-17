package main

import (
	"context"
	"fmt"
	"time"

	"github.com/atmetz/therapy-scheduler/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerPlatformsCreate(params parameters) (database.Platform, error) {

	platform, err := cfg.db.CreatePlatform(context.Background(), database.CreatePlatformParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return database.Platform{}, fmt.Errorf("couldn't create platform %s", err)
	}

	return platform, nil

}

func (cfg *apiConfig) getPlatformList() ([]string, error) {

	platforms, err := cfg.db.GetPlatforms(context.Background())
	if err != nil {
		return []string{}, err
	}
	var platformList []string

	for _, platform := range platforms {
		platformList = append(platformList, platform.Name)
	}

	return platformList, nil
}

func (cfg *apiConfig) getPlatformID(name string) (uuid.UUID, error) {
	platform, err := cfg.db.GetPlatformByName(context.Background(), name)
	if err != nil {
		return uuid.Nil, err
	}

	return platform.ID, nil
}
