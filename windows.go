package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (cfg *apiConfig) platformWindow(a fyne.App, actionLabel *widget.Label) {
	platformWindow := a.NewWindow("Add Platform")
	platformWindow.Resize(fyne.NewSize(300, 300))
	platformWindow.CenterOnScreen()
	platformName := widget.NewEntry()
	var actionMessage string
	errorMessage := ""
	errorLabel := widget.NewLabel(errorMessage)

	// Create platform entry form
	platformForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: platformName},
		},
		// Submit button
		OnSubmit: func() {
			_, err := cfg.handlerPlatformsCreate(parameters{
				Name: platformName.Text,
			})
			if err != nil {
				errorMessage = fmt.Sprintf("error creating platform: %s\n", err)
				errorLabel.SetText(errorMessage)
			} else {
				actionMessage = fmt.Sprintf("Platform Created: %s", platformName.Text)
				actionLabel.SetText(actionMessage)
				platformWindow.Close()
			}
		},
		OnCancel: func() {
			platformWindow.Close()
		},
	}
	platformWindow.SetContent(container.NewVBox(
		platformForm,
		errorLabel,
	))
	platformWindow.Show()

}

func (cfg *apiConfig) clientWindow(a fyne.App, actionLabel *widget.Label) {
	clientWindow := a.NewWindow("Add Client")
	clientWindow.Resize(fyne.NewSize(300, 300))
	clientWindow.CenterOnScreen()
	clientName := widget.NewEntry()
	clientEmail := widget.NewEntry()
	clientPhone := widget.NewEntry()
	errorMessage := ""
	errorLabel := widget.NewLabel(errorMessage)

	options, err := cfg.getPlatformList()
	if err != nil {
		errorMessage = fmt.Sprintf("%s", err)
		errorLabel.SetText(errorMessage)
	}
	var actionMessage string
	var selectedValue string
	selectWidget := widget.NewSelect(options, func(value string) {
		selectedValue = value
	})
	// New Client Form
	clientForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: clientName},
			{Text: "Email", Widget: clientEmail},
			{Text: "Phone Number", Widget: clientPhone},
			{Text: "Platform", Widget: selectWidget},
		},
		OnSubmit: func() {
			// Add Platform

			platformID, err := cfg.getPlatformID(selectedValue)
			if err != nil {
				errorMessage = fmt.Sprintf("%s", err)
				errorLabel.SetText(errorMessage)
			} else {

				client, err := cfg.handlerClientsCreate(parameters{
					Name:       clientName.Text,
					Email:      clientEmail.Text,
					Phone:      clientPhone.Text,
					ProviderID: cfg.currentUser.ID,
					PlatformID: platformID,
				})
				if err != nil {
					errorMessage = fmt.Sprintf("%s", err)
					errorLabel.SetText(errorMessage)
				} else {
					clientWindow.Close()
					actionMessage = fmt.Sprintf("Client Created: %s", client.Name)
					actionLabel.SetText(actionMessage)
				}
			}
		},
		OnCancel: func() {
			clientWindow.Close()
		},
	}

	clientWindow.SetContent(clientForm)
	clientWindow.Show()
}

func (cfg *apiConfig) providerLoginWindow(a fyne.App, welcomeLabel *widget.Label) {
	providersWindow := a.NewWindow("Provider Login")
	providersWindow.Resize(fyne.NewSize(300, 300))
	providersWindow.CenterOnScreen()
	providerPassword := widget.NewEntry()
	providerEmail := widget.NewEntry()
	errorMessage := ""
	errorLabel := widget.NewLabel(errorMessage)

	// Provider Login Form
	loginForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Email", Widget: providerEmail},
			{Text: "Password", Widget: providerPassword},
		},
		OnSubmit: func() {
			user, err := cfg.handlerProviderLogin(parameters{
				Email:    providerEmail.Text,
				Password: providerPassword.Text,
			})
			if err != nil {
				errorMessage = fmt.Sprintf("Login credentials incorrect: %s", err)
				errorLabel.SetText(errorMessage)
			} else {
				providersWindow.Close()
				cfg.currentUser = user
				welcomeMessage := fmt.Sprintf("Welcome %s", cfg.currentUser.Name)
				welcomeLabel.SetText(welcomeMessage)
			}
		},
		OnCancel: func() {
			providersWindow.Close()
		},
	}

	providersWindow.SetContent(container.NewVBox(
		loginForm,
		errorLabel,
	))
	providersWindow.Show()
}

func (cfg *apiConfig) newProviderWindow(a fyne.App, welcomeLabel *widget.Label) {
	providersWindow := a.NewWindow("Register Provider")
	providersWindow.Resize(fyne.NewSize(300, 300))
	providersWindow.CenterOnScreen()
	providerName := widget.NewEntry()
	providerPassword := widget.NewEntry()
	providerEmail := widget.NewEntry()
	providerPhone := widget.NewEntry()
	errorMessage := ""
	errorLabel := widget.NewLabel(errorMessage)

	// Create new provider form
	providerForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: providerName},
			{Text: "Email", Widget: providerEmail},
			{Text: "Phone", Widget: providerPhone},
			{Text: "Password", Widget: providerPassword},
		},
		OnSubmit: func() {
			user, err := cfg.handlerProvidersCreate(parameters{
				Name:     providerName.Text,
				Email:    providerEmail.Text,
				Phone:    providerPhone.Text,
				Password: providerPassword.Text,
			})
			if err != nil {

				errorMessage = fmt.Sprintf("%s", err)
				errorLabel.SetText(errorMessage)
			} else {
				providersWindow.Close()
				cfg.currentUser = user
				welcomeMessage := fmt.Sprintf("Welcome %s", cfg.currentUser.Name)
				welcomeLabel.SetText(welcomeMessage)
			}
		},
		OnCancel: func() {
			providersWindow.Close()
		},
	}

	providersWindow.SetContent(providerForm)
	providersWindow.Show()
}
