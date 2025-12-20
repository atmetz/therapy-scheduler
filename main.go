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
	db          database.Queries
	currentUser database.Provider
}

type parameters struct {
	Name              string
	Email             string
	Phone             string
	Password          string
	Frequency         string
	ProviderID        uuid.UUID
	PlatformID        uuid.UUID
	StartDate         time.Time
	SessionsAvailable int64
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
	currentUser := database.Provider{
		Name: "User",
	}

	cfg := apiConfig{
		db:          *dbQueries,
		currentUser: currentUser,
	}

	a := app.New()
	w := a.NewWindow("Therapy Scheduler")
	w.Resize(fyne.NewSize(1000, 1000))
	w.CenterOnScreen()

	welcomeMessage := fmt.Sprintf("Welcome %s", cfg.currentUser.Name)
	welcomeLabel := widget.NewLabel(welcomeMessage)
	actionMessage := ""
	actionLabel := widget.NewLabel(actionMessage)

	w.SetContent(container.NewVBox(
		welcomeLabel,
		actionLabel,
		// Create new providers window
		widget.NewButton("Register Provider", func() {
			cfg.newProviderWindow(a, welcomeLabel)
		}),
		// Provider Login
		widget.NewButton("Provider Login", func() {
			cfg.providerLoginWindow(a, welcomeLabel)
		}),
		// Add new client
		widget.NewButton("Add Client", func() {
			cfg.clientWindow(a, actionLabel)
		}),
		// Add new platform
		widget.NewButton("Add Platform", func() {
			cfg.platformWindow(a, actionLabel)
		}),
	))

	w.Show()
	a.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}
