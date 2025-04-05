package data_interface

import (
	"encoding/csv"
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

	fd := FoodData{name, enabled, last_date_tried, likes}	
	return &fd
}

func (fd *FoodData) ToStringSlice() []string {
	return []string{fd.name, strconv.FormatBool(fd.enabled), fd.last_date_tried.String(), strconv.FormatBool(fd.likes)}

} 

func Print_csv() {
	_, serr := os.Stat(file_name)
	
	if serr != nil {
		fmt.Printf("Error stat finding file: %v", serr)
		return
	}
	file, err := os.OpenFile(file_name, os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
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
	fd := newFoodData(food_name, enabled, last_date_tried, likes)
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Error opening file: %v", err)
		return false
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return false
	}

	written := false

	for i, record := range records {
		if record[0] == food_name {
			records[i] = fd.ToStringSlice()
			written = true
		}
	}

	if written == true {
		file.Seek(0, 0)
		writer := csv.NewWriter(file)
		defer writer.Flush()
		err = writer.WriteAll(records)

		if err != nil {
			fmt.Println("Error writing to CSV:", err)
			return false
		}
		return true
	}

	return false
}
