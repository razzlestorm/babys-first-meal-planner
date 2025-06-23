package models

import (
	surrealdb "github.com/surrealdb/surrealdb.go"
	"github.com/surrealdb/surrealdb.go/pkg/models"
)


type FoodData struct {
	ID            *models.RecordID `json:"id,omitempty"`
	Name          string           `json:"name"`
	Enabled       bool		`json:"enabled"`
	LastDateTried time.Time `json:"date_tried"`
	Likes         bool `json:"likes"`
}

type FoodDataModel struct {
	db *surrealdb.DB
}


func (m *FoodDataModel) Insert(name string, enabled likes bool, lastDateTried time.Time) (int, error) {
	createdFoods, err := surrealdb.Create[[]FoodData](db, models.Table("food_data"), FoodData{
		Name:    name,
		Enabled: enabled,
		Likes: likes,
		LastDateTried: lastDateTried, 
	})

	
	if err != nil {
		return 0, err
	}

	food1 := FirstOrNil[FoodData](*createdFoods)
	if food1 == nil {
		panic("No food data returned")
	}
	return int(food1.id), nil
}
