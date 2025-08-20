package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/razzlestorm/babys-first-meal-planner/cmd/calendar"
	"github.com/razzlestorm/babys-first-meal-planner/internal/models"
)

type templateData struct {
	Foods  []models.FoodData
	Config *calendar.CalendarConfig
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

func makeRangeInt(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = i
	}
	return r
}

// Initialize a template.FuncMap object and store it in a global variable. This is
// essentially a string-keyed map which acts as a lookup between the names of our
// custom template functions and the functions themselves.
var functions = template.FuncMap{
	"humanDate":    humanDate,
	"makeRangeInt": makeRangeInt,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./cmd/server/ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("./cmd/server/ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./cmd/server/ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
