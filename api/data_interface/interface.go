package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"github.com/razzlestorm/babys-first-meal-planner/api/data_interface/interface"
	"github.com/razzlestorm/babys-first-meal-planner/data"
)

func main() {
	file, err := os.Open("data/food_db.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	for _, record := range records {
		fmt.Println(record)
	}
}
