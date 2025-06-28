package models

import (
	"time"
	"database/sql"
	"errors"
)

type FoodData struct {
	ID            int `json:"id,omitempty"`
	Name          string           `json:"name"`
}

type MealPlannerModel struct {
	DB *sql.DB
}


func (m *MealPlannerModel) InsertFood(name string) (int, error) {
	stmt := `INSERT INTO Foods (name) VALUES(?)`
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
	stmt := `SELECT id, name FROM foods WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	var fd FoodData

	err := row.Scan(&fd.ID, &fd.Name, &fd.Enabled, &fd.Likes, &fd.LastDateTried)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return FoodData{}, ErrNoRecord
		} else {
			return FoodData{}, err
		}
	}

	return fd, nil
}

