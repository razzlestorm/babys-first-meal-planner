package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/razzlestorm/babys-first-meal-planner/cmd/calendar"
	"github.com/razzlestorm/babys-first-meal-planner/internal/models"
)

type application struct {
	logger         *slog.Logger
	templateCache  map[string]*template.Template
	calendarConfig *calendar.CalendarConfig
	planner        *models.MealPlannerModel
}

func openDB(user, pass, dbName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s?parseTime=true", user, pass, dbName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("PORT")

	db, err := openDB(user, pass, dbName)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Check for user being logged in here and use their calendar
	calendarConfig, err := calendar.NewCalendarConfig(30, 3, 3)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// eventually serve our application.routes(), where we will list the various page routes to go to
	// For now though, we're just going to have it on the main page
	// create infinte loop
	// Run manager
	// As user logs in , populate manager.sessions with userSessions with a timeout
	// Return to a saved session, or start a new one if there wasn't a previous session
	// Load in the calendar config and the latest calendar that they were working on (or new)
	app := &application{
		logger:         logger,
		templateCache:  templateCache,
		calendarConfig: calendarConfig,
		planner:        &models.MealPlannerModel{DB: db},
	}

	logger.Info("Starting server", "port", port)
	err = http.ListenAndServe(port, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
