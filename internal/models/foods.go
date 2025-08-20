package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type Category int

const (
	Fruit Category = iota
	Vegetable
	Grain
	Protein
	Dairy
)

type FoodData struct {
	ID       int      `json:"id,omitempty"`
	Name     string   `json:"name"`
	Category Category `json:"category"`
}

func (f FoodData) GetName() string {
	return f.Name
}

func (f FoodData) GetCategory() Category {
	return f.Category
}

func (m *MealPlannerModel) InsertFood(name string, category Category) (int, error) {
	stmt := `INSERT INTO Foods (name, category) VALUES(?, ?)`
	result, err := m.DB.Exec(stmt, name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *MealPlannerModel) GetFood(id int) (FoodData, error) {
	stmt := `SELECT id, name, category FROM foods WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	var fd FoodData

	err := row.Scan(&fd.ID, &fd.Name, &fd.Category)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return FoodData{}, ErrNoRecord
		} else {
			return FoodData{}, err
		}
	}

	return fd, nil
}

func (m *MealPlannerModel) GetAllFoods() ([]FoodData, error) {
	stmt := `SELECT id, name, category FROM foods`

	rows, err := m.DB.Query(stmt)

	defer rows.Close()

	var fd []FoodData

	for rows.Next() {

		var f FoodData

		var category string
		if err = rows.Scan(&f.ID, &f.Name, &category); err != nil {
			return nil, err
		}

		// categories are stored as strings in DB
		switch category {
		case "fruit":
			f.Category = Fruit
		case "vegetable":
			f.Category = Vegetable
		case "grain":
			f.Category = Grain
		case "protein":
			f.Category = Protein
		case "dairy":
			f.Category = Dairy
		default:
			return nil, fmt.Errorf("unknown category: %s", category)
		}

		fd = append(fd, f)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return fd, nil
}
