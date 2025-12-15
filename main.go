package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atmetz/therapy-scheduler/internal/database"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	db database.Queries
}

type provider struct {
	ID          uuid.UUID
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Password    string
	PhoneNumber string
	Email       string
}

func main() {

	godotenv.Load(".env")

	pathToDB := os.Getenv("DB_PATH")
	if pathToDB == "" {
		log.Fatal("DB_PATH must be set")
	}

	db, err := sql.Open("sqlite3", pathToDB)
	if err != nil {
		log.Fatalf("Couldn't connect to database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	cfg := apiConfig{
		db: *dbQueries,
	}

	a := app.New()
	w := a.NewWindow("Therapy Scheduler")
	w.Resize(fyne.NewSize(1000, 1000))
	w.CenterOnScreen()

	currentUser := provider{
		Name: "User",
	}
	welcomeMessage := fmt.Sprintf("Welcome %s", currentUser.Name)
	hello := widget.NewLabel(welcomeMessage)

	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Register Provider", func() {
			// Create new providers window
			providersWindow := a.NewWindow("Register Provider")
			providersWindow.Resize(fyne.NewSize(300, 300))
			providersWindow.CenterOnScreen()
			providerName := widget.NewEntry()
			providerPassword := widget.NewEntry()
			providerEmail := widget.NewEntry()
			providerPhone := widget.NewEntry()

			// Create new provider form
			providerForm := &widget.Form{
				Items: []*widget.FormItem{
					{Text: "Name", Widget: providerName},
					{Text: "Email", Widget: providerEmail},
					{Text: "Phone", Widget: providerPhone},
					{Text: "Password", Widget: providerPassword},
				},
				OnSubmit: func() {
					currentUser, err := cfg.handlerProvidersCreate(parameters{
						Name:     providerName.Text,
						Email:    providerEmail.Text,
						Phone:    providerPhone.Text,
						Password: providerPassword.Text,
					})
					if err != nil {
						log.Printf("%v\n", err)
					}
					providersWindow.Close()
					welcomeMessage = fmt.Sprintf("Welcome %s", currentUser.Name)
					hello.SetText(welcomeMessage)
				},
				OnCancel: func() {
					providersWindow.Close()
				},
			}

			providersWindow.SetContent(providerForm)
			providersWindow.Show()
		}),
		// Provider Login Widget
		widget.NewButton("Provider Login", func() {
			providersWindow := a.NewWindow("Provider Login")
			providersWindow.Resize(fyne.NewSize(300, 300))
			providersWindow.CenterOnScreen()
			providerPassword := widget.NewEntry()
			providerEmail := widget.NewEntry()
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
						fmt.Println(err)
					} else {
						providersWindow.Close()
						welcomeMessage = fmt.Sprintf("Welcome %s", user.Name)
						hello.SetText(welcomeMessage)
					}
				},
				OnCancel: func() {
					providersWindow.Close()
				},
			}

			providersWindow.SetContent(loginForm)
			providersWindow.Show()
		}),
	))

	w.Show()
	a.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
