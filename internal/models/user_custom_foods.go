package models

import (
	"time"
	"database/sql"
	"errors"
)

type UserCustomFoodData struct {
	ID            int `json:"id,omitempty"`
	UserID		string
	Name          string           `json:"name"`
}


func (m *MealPlannerModel) InsertCustomFood(userId, name string) (int, error) {
	stmt := `INSERT INTO UserCustomFoods (user_id, name) VALUES(?, ?)`
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
	stmt := `SELECT id, name FROM UserCustomFoods WHERE id = ?`

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

