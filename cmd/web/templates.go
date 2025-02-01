package main

import (
	"html/template"
	"net/http"
)

type templateData struct {
	Talks    []Talk
	Speakers []Speaker
	Events   []Event
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	// Create template set
	ts, err := template.ParseFiles("./ui/html/pages/"+page, "./ui/html/pages/base.html", "./ui/html/partials/nav.html")
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
