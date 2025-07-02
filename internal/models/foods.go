package models

import (
	"database/sql"
	"errors"
	"time"
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

func (f FoodData) GetName() {
	return f.Name
}

func (f FoodData) GetCategory() {
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
