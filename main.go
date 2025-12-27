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

// table example
var data = [][]string{
	{"", ""},
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

	//create table to display client list
	var columnHeaders = []string{"Client Name", "Sessions"}
	table := widget.NewTableWithHeaders(
		// Length callback: returns the number of rows and columns
		func() (int, int) {
			return len(data), len(data[0])
		},
		// CreateCell callback: returns a new template object for a cell
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		// UpdateCell callback: is called to apply data to a cell template
		func(id widget.TableCellID, object fyne.CanvasObject) {
			label := object.(*widget.Label)
			label.SetText(data[id.Row][id.Col])
		},
	)

	table.SetColumnWidth(0, 100)
	table.SetColumnWidth(1, 10)

	table.UpdateHeader = func(id widget.TableCellID, obj fyne.CanvasObject) {
		if id.Col >= 0 && id.Col < len(columnHeaders) {
			obj.(*widget.Label).SetText(columnHeaders[id.Col])
		}
	}

	// File menu
	quitItem := fyne.NewMenuItem("Quit", func() {
		a.Quit()
	})
	fileMenu := fyne.NewMenu("File", quitItem)

	// Provider menu
	providerItem := fyne.NewMenuItem("Provider Login", func() {
		cfg.providerLoginWindow(a, welcomeLabel, actionLabel, table)
	})
	providerRegisterItem := fyne.NewMenuItem("New Provider", func() {
		cfg.newProviderWindow(a, welcomeLabel)
	})
	providerMenu := fyne.NewMenu("Provider",
		providerItem,
		providerRegisterItem)

	// Client menu
	clientItem := fyne.NewMenuItem("Add Client", func() {
		cfg.clientWindow(a, actionLabel)
	})
	seeClientItem := fyne.NewMenuItem("Show Clients", func() {
		cfg.updateClientTable(actionLabel, table)
	})
	clientMenu := fyne.NewMenu("Client",
		clientItem,
		seeClientItem)

	// Platform menu
	platformItem := fyne.NewMenuItem("Add Platform", func() {
		cfg.platformWindow(a, actionLabel)
	})
	platformMenu := fyne.NewMenu("Platform", platformItem)

	// Main menu
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		providerMenu,
		clientMenu,
		platformMenu,
	)

	w.SetMainMenu(mainMenu)

	content := container.NewBorder(
		container.NewVBox(
			welcomeLabel,
			actionLabel,
		),
		nil,
		nil,
		nil,
		table)
	w.SetContent(content)

	w.Show()
	a.Run()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}

func (cfg *apiConfig) updateClientTable(actionLabel *widget.Label, table *widget.Table) {
	clientList, err := cfg.handlerShowClients()
	if err != nil {
		actionMessage := fmt.Sprintf("error: %s", err)
		actionLabel.SetText(actionMessage)
		return
	}
	data, err = cfg.getClientsData(clientList)
	if err != nil {
		actionMessage := fmt.Sprintf("error: %s", err)
		actionLabel.SetText(actionMessage)
		data = [][]string{
			{"", ""},
		}
		table.Refresh()
		return
	}
	table.Refresh()
}
