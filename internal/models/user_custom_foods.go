package models

import (
	"database/sql"
	"errors"
	"time"
)

type UserCustomFoodData struct {
	ID       int `json:"id,omitempty"`
	UserID   string
	Name     string   `json:"name"`
	Category Category `json:"category"`
}

func (m *MealPlannerModel) InsertCustomFood(userId, name string, category Category) (int, error) {
	stmt := `INSERT INTO UserCustomFoods (user_id, name, category) VALUES(?, ?, ?)`
	result, err := m.DB.Exec(stmt, userId, name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *MealPlannerModel) GetCustomFood(id int) (FoodData, error) {
	stmt := `SELECT id, name, category FROM UserCustomFoods WHERE user_id = ?`

	row := m.DB.QueryRow(stmt, id)

	var fd UserCustomFoodData

	err := row.Scan(&fd.ID, &fd.Name, &fd.Category)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return UserCustomFoodData{}, ErrNoRecord
		} else {
			return UserCustomFoodData{}, err
		}
	}

	return fd, nil
}
