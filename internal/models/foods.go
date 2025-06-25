package models

import (
	"time"
	"database/sql"
	"errors"
)

type FoodData struct {
	ID            int `json:"id,omitempty"`
	Name          string           `json:"name"`
	Enabled       bool             `json:"enabled"`
	Likes         bool             `json:"likes"`
	LastDateTried time.Time        `json:"date_tried"`
}

type FoodDataModel struct {
	DB *sql.DB
}


func (m *FoodDataModel) Insert(name string, enabled, likes bool, lastDateTried time.Time) (int, error) {
	stmt := `INSERT INTO foods (name, enabled, likes, last_date_tried) VALUES(?, ?, ?, ?)`
	result, err := m.DB.Exec(stmt, name, enabled, likes, lastDateTried)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}


func (m *FoodDataModel) Get(id int) (FoodData, error) {
	stmt := `SELECT id, name, enabled, likes, last_date_tried FROM foods WHERE id = ?`

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

