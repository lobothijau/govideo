package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)

	mux.HandleFunc("GET /speaker/create", app.speakerCreate)
	mux.HandleFunc("POST /speaker/create", app.speakerCreatePost)

	mux.HandleFunc("GET /event/create", app.eventCreate)
	mux.HandleFunc("POST /event/create", app.eventCreatePost)

	mux.HandleFunc("GET /talks/create", app.talkCreate)
	mux.HandleFunc("POST /talks/create", app.talkCreatePost)
	mux.HandleFunc("GET /talks/{id}", app.talkDetail)
	mux.HandleFunc("GET /talks", app.talksView)

	mux.HandleFunc("GET /speakers/{id}", app.speakerDetail)

	return mux
}
