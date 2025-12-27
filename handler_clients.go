package main

import (
	"context"
	"errors"
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
		Frequency:   params.Frequency,
		PlatformID:  params.PlatformID,
		ProviderID:  params.ProviderID,
		StartDate:   params.StartDate,
		EndDate:     time.Time{},
	})
	if err != nil {
		return database.Client{}, fmt.Errorf("couldn't create client %s", err)
	}

	return user, nil

}

func (cfg *apiConfig) handlerShowClients() ([]database.Client, error) {
	clients, err := cfg.db.GetClientsByProvider(context.Background(), cfg.currentUser.ID)
	if err != nil {
		return []database.Client{}, err
	}

	return clients, nil
}

func (cfg *apiConfig) getClientsData(clients []database.Client) ([][]string, error) {

	var clientData [][]string

	if len(clients) == 0 {
		return [][]string{}, errors.New("Provider does not have any clients")
	}

	for _, client := range clients {
		//clientData[i] = append(clientData[i], client.Name)
		switch client.Frequency {
		case "Weekly":
			clientData = append(clientData, []string{client.Name, "4"})
		case "Every other week":
			clientData = append(clientData, []string{client.Name, "2"})
		case "Once a month":
			clientData = append(clientData, []string{client.Name, "1"})
		case "Every other month":
			clientData = append(clientData, []string{client.Name, "0.5"})
		case "Schedules session after each session":
			clientData = append(clientData, []string{client.Name, "1"})
		case "No appointments scheduled":
			clientData = append(clientData, []string{client.Name, "0"})
		}
	}
	return clientData, nil
}
