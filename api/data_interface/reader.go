package data_interface

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Print_csv() {
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
