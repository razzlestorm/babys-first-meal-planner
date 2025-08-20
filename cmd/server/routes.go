package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	//init new servemux and point to home func to handle "/" URL pattern
	mux := http.NewServeMux()

	// Create a file server which servbes files out of "./ui/static". Path relative to project root
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// mux Handle registers file server as handler for all URL paths beginning with "/static/"
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /foods", app.foods)

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
