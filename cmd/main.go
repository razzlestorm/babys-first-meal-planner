package main

import (
	"time"
	"github.com/razzlestorm/babys-first-meal-planner/api/data_interface"
)
func main() {
	
	// data_interface.Print_csv()
	test := data_interface.Write_row("Egg", true, time.Now(), true)
	println(test)
}
