package data_interface

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

var file_name = "data/food_db.csv";

type FoodData struct {
	name string
	enabled bool
	last_date_tried time.Time
	likes bool
}

func newFoodData(name string, enabled bool, last_date_tried time.Time, likes bool) *FoodData {

	
}

func Print_csv() {
	file, err := os.Open(file_name)
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

// chage this to write_rows when struct is created, like just take in []string
func Write_row(food_name string, enabled bool, last_date_tried time.Time, likes bool) bool {
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return false
	}

	for i, record := range records {
		if record[0] == food_name {
			records[i] = []string{food_name, strconv.FormatBool(enabled), last_date_tried.String(), strconv.FormatBool(likes)}
			return true
		}
	}

	return false
}
