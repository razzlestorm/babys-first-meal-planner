package main

import (
	"fmt"
	surrealdb "github.com/surrealdb/surrealdb.go"
	"github.com/surrealdb/surrealdb.go/pkg/models"
	"time"
	// "github.com/razzlestorm/babys-first-meal-planner/api/data_interface"
)

type FoodData struct {
	ID            *models.RecordID `json:"id,omitempty"`
	Name          string           `json:"name"`
	Enabled       bool             `json:"enabled"`
	LastDateTried time.Time        `json:"date_tried"`
	Likes         bool             `json:"likes"`
}

func FirstOrNil[T any](slice []T) *T {
	if len(slice) == 0 {
		return nil
	}
	return &slice[0]
}

func main() {

	// Connect to SurrealDB
	db, err := surrealdb.New("ws://localhost:8000")
	if err != nil {
		panic(err)
	}

	// Set the namespace and database
	if err = db.Use("testNS", "testDB"); err != nil {
		panic(err)
	}

	// Sign in to authentication `db`
	authData := &surrealdb.Auth{
		Username: "root", // use your setup username
		Password: "root", // use your setup password
	}
	token, err := db.SignIn(authData)
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
	// data_interface.Print_csv()
	// test := data_interface.Write_row("Egg", true, time.Now(), true)
	// println(test)// Or use structs
	createdFoods, err := surrealdb.Create[[]FoodData](db, models.Table("food_data"), FoodData{
		Name:    "Cheese",
		Enabled: true,
	})
	if err != nil {
		panic(err)
	}

	food1 := FirstOrNil[FoodData](*createdFoods)
	if food1 == nil {
		panic("No food data returned")
	}
	fmt.Printf("Created food data with a struct: %+v\n", *food1)

	// Get entry by Record ID
	selectedFood, err := surrealdb.Select[FoodData, models.RecordID](db, *food1.ID)
	if err != nil {
		panic(err)
	}

	if selectedFood == nil {
		panic("No food found with that ID")
	}

	fmt.Printf("Selected a food by record id: %+v\n", selectedFood)
	// Or retrieve the entire table
	//foods, err := surrealdb.Select[[]FoodData, models.Table](db, models.Table("food_data"))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Selected all in food_data table: %+v\n", foods)
}
