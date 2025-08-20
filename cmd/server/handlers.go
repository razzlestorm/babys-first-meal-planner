package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	data := templateData{Foods: nil, Config: app.calendarConfig}

	app.render(w, r, http.StatusOK, "home.tmpl", data)
}

func (app *application) foods(w http.ResponseWriter, r *http.Request) {

	foods, err := app.planner.GetAllFoods()

	if err != nil {
		app.serverError(w, r, err)
		return
	}
	data := templateData{Foods: foods, Config: app.calendarConfig}

	app.render(w, r, http.StatusOK, "foods.tmpl", data)
}
