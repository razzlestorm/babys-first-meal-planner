package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log/slog"
	"new/http"
	"os"
	"time"

	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/razzlestorm/babys-first-meal-planner/internal/models"
)

type application struct {
	logger        *slog.Logger
	templateCache map[string]*template.Template
	formDecoder   *form.Decoder
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

func home() {
	
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

	db, err := openDB(user, pass, dbName)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	/*
		planner := models.MealPlannerModel{DB: db}

		foodID, err := planner.InsertFood("Cheese")

		if err != nil {
			panic(err)
		}

		fmt.Printf("Created food data with an id: %+d\n", foodID)

		// Get entry by Record ID
		selectedFood, err := fdm.Get(foodID)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Selected a food by record id: %+v\n", selectedFood)
	*/
	// eventually serve our application.routes(), where we will list the various page routes to go to
	// For now though, we're just going to have it on the main page
	// create infinte loop
	// Run manager
	// As user logs in , populate manager.sessions with userSessions with a timeout
	// Return to a saved session, or start a new one if there wasn't a previous session
}
