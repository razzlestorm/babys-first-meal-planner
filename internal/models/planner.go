package models

import (
	"time"
	"database/sql"
	"errors"
)

type MealPlannerModel struct {
	DB *sql.DB
}



type UserData struct {
	UserID		string
	Email	        string
}


// TODO: Separate this out into its own handler, not MealPlannerModel
func (m *MealPlannerModel) InsertUser(userId, email string) (string, error) {
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

	var ud UserData

	err := row.Scan(&ud.UserID, &ud.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return ud.UserID, nil
}

