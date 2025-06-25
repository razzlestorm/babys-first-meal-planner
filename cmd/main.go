package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"

	"github.com/razzlestorm/babys-first-meal-planner/internal/models"
)


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

	db, err := openDB(user, pass, dbName)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()


	fdm := models.FoodDataModel{DB: db}

	foodID, err := fdm.Insert("Cheese", true, true, time.Now())

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
}
