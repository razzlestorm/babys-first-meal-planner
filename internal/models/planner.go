package models

import (
	"database/sql"
	"errors"
)

type MealPlannerModel struct {
	DB *sql.DB
}

type UserData struct {
	UserID string
	Email  string
}

// TODO: Separate this out into its own handler, not MealPlannerModel
// TODO: Also update the id to be uuids
func (m *MealPlannerModel) InsertUser(userId, email string) (int64, error) {
	stmt := `INSERT INTO Users (userId, email) VALUES(?, ?)`
	result, err := m.DB.Exec(stmt, userId, email)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *MealPlannerModel) GetUserId(id string) (string, error) {
	stmt := `SELECT user_id FROM Users WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	var userID string

	err := row.Scan(&userID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrNoRecord
		} else {
			return "", err
		}
	}

	return userID, nil
}
