package main

import (
	"html/template"
	"net/http"

	"govideo.dev/internal/models"
)

type templateData struct {
	Talks    []Talk
	Speakers []models.Speaker
	Events   []models.Event
}

func (app *application) render(w http.ResponseWriter, r *http.Request, page string, data templateData) {
	ts, err := template.ParseFiles("./ui/html/pages/"+page, "./ui/html/pages/base.html", "./ui/html/partials/nav.html")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
