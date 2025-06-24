package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/razzlestorm/babys-first-meal-planner/internal/models"
	surrealdb "github.com/surrealdb/surrealdb.go"
)

func openDB(auth *surrealdb.Auth, dsn, namespace, db_name string) (*surrealdb.DB, error) {
	db, err := surrealdb.New(dsn)
	if err != nil {
		panic(err)
	}

	// Set the namespace and database
	if err = db.Use(namespace, db_name); err != nil {
		panic(err)
	}

	// Sign in to authentication `db`
	token, err := db.SignIn(auth)
	if err != nil {
		panic(err)
	}

	// Check token validity. This is not necessary if you called `SignIn` before. This authenticates the `db` instance too if sign in was
	// not previously called
	if err := db.Authenticate(token); err != nil {
		panic(err)
	}

	// And we can later on invalidate the token if desired
	defer func(token string) {
		if err := db.Invalidate(); err != nil {
			panic(err)
		}
	}(token)

	return db, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	authData := &surrealdb.Auth{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
	}
	// Connect to SurrealDB
	db, err := openDB(authData, "ws://localhost:8000", os.Getenv("DB_NAMESPACE"), os.Getenv("DB_NAME"))
	defer db.Close()

	fdm := models.FoodDataModel{DB: db}

	// testing that this gets created
	// SurrealDB doesn't support time.Now() apparently. It can't convert time.Times into its own DT format.
	// Consider switching DBs?
	createdFood, err := fdm.Insert("Cheese", true, true, time.Now())

	if err != nil {
		panic(err)
	}

	fmt.Printf("Created food data with a struct: %+v\n", *createdFood)

	// Get entry by Record ID
	selectedFood, err := fdm.Get(createdFood.ID.String())
	if err != nil {
		panic(err)
	}

	if selectedFood == nil {
		panic("No food found with that ID")
	}

	fmt.Printf("Selected a food by record id: %+v\n", selectedFood)
}
