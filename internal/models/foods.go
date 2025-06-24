package models

import (
	surrealdb "github.com/surrealdb/surrealdb.go"
	"fmt"
	"github.com/surrealdb/surrealdb.go/pkg/models"
	"time"
)

type FoodData struct {
	ID            *models.RecordID `json:"id,omitempty"`
	Name          string           `json:"name"`
	Enabled       bool             `json:"enabled"`
	Likes         bool             `json:"likes"`
	LastDateTried time.Time        `json:"date_tried"`
}

type FoodDataModel struct {
	DB *surrealdb.DB
}

func FirstOrNil[T any](slice []T) *T {
	if len(slice) == 0 {
		return nil
	}
	return &slice[0]
}

func (m *FoodDataModel) Insert(name string, enabled, likes bool, lastDateTried time.Time) (*FoodData, error) {
	fmt.Printf("%+v\n", FoodData{
		Name:          name,
		Enabled:       enabled,
		Likes:         likes,
		LastDateTried: lastDateTried,
	})
	createdFoods, err := surrealdb.Create[[]FoodData](m.DB, models.Table("food_data"), FoodData{
		Name:          name,
		Enabled:       enabled,
		Likes:         likes,
		LastDateTried: lastDateTried,
	})

	if err != nil {
		return nil, err
	}

	food := FirstOrNil(*createdFoods)
	if food == nil {
		panic("No food data returned")
	}
	return food, nil
}

func (m *FoodDataModel) Get(id string) (*FoodData, error) {
	food, err := surrealdb.Select[FoodData](m.DB, id)

	if err != nil {
		return nil, err
	}

	return food, nil

}
