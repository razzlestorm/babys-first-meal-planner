package calendar

import (
	_ "math/rand"
	"fmt"
	"github.com/razzlestorm/babys-first-meal-planner/internal/models"
)

// create an interface for FoodData and UserCustomFoodData
type Food interface {
	GetName() string
	GetCategory() models.Category
}

// Figure out what this wants to look like later
type Job struct{}

type CalendarConfig struct {
	jobChan      chan Job
	days         int
	mealsPerDay  int
	foodsPerMeal int
	foods        []Food
}


func NewCalendarConfig(days int, mpd int, fpm int) (*CalendarConfig, error) {

	if days <= 0 || mpd <= 0 || fpm <= 0 {
		return nil, fmt.Errorf("all arguments must be greater than 0: got days=%d, mealsPerDay=%d, foodsPerMeal=%d", days, mpd, fpm)
	}

	c := CalendarConfig{
		jobChan: make(chan Job),
		days: days,
		mealsPerDay: mpd,
		foodsPerMeal: fpm,
		// Figure out how/when we want to add foods, maybe in another function
		foods: []Food{},
	}
	return &c, nil
}


type Calendar struct {
	mapping map[int]map[int]string
}

// The buffer is the amount of "buffer" meals before a food can be repeated.
// 0 = a food can be (but isn't necessarily)repeated every meal, 1 = a food can be repeated every other meal, etc.
// If you wanted no repeats, offset = c.days
func (c *Calendar) randomize(buffer int, config CalendarConfig) {
	/*
	   have two sets, one = foods, one = used := map[name]buffer (then count down, add it back to possible foods when it reaches 0) If buffer = 0, don't need to do this
	   1. Loop over days
	   2. Loop over meals per day
	   3. select random foods from available foods
	   4. save calendar to user session, will look like:
	   1[1] = carrots, steak, grapes
	   1[2] = lemon, salmon, lettuce

	for day := range c.days {
		
	}
	*/

}
