package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := templateData{}

	// Simulate talks data
	talks, err := app.talks.GetLatest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Simulate speakers data
	speakers, err := app.speakers.GetActive()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Simulate events data
	events, err := app.events.GetLatest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data.Talks = talks
	data.Speakers = speakers
	data.Events = events

	app.render(w, r, "home.html", data)
}

func (app *application) videoView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Video View!"))
}

func (app *application) videoDetail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/video/"):])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Display video detail with ID  of %d", id)
	w.Write([]byte(msg))
}

func (app *application) videoCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Video Create!"))
}
