package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"govideo.dev/internal/models"
	"govideo.dev/internal/view"
)

type templateData struct {
	Talk         *models.Talk
	Speaker      *models.Speaker
	Talks        []models.Talk
	RelatedTalks []models.Talk
	Speakers     []models.Speaker
	Events       []models.Event
	CurrentPage  int
	TotalPages   int
	HasNext      bool
	HasPrev      bool
	NextPage     int
	PrevPage     int
	Pages        []int
}

// cache to hold our templates
var templateCache map[string]*template.Template

func init() {
	templateCache = make(map[string]*template.Template)
}

// initTemplates initializes our template cache
func (app *application) initTemplates() error {
	// Define our template functions
	functions := template.FuncMap{
		"formatDateRange":  view.FormatDateRange,
		"avatarURL":        view.AvatarURL,
		"eventBannerURL":   view.EventBannerURL,
		"talkThumbnailURL": view.TalkThumbnailURL,
	}

	// Get all page templates
	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return err
	}

	// Loop through each page
	for _, page := range pages {
		name := filepath.Base(page)

		// Create a template set with our custom functions
		ts, err := template.New(name).Funcs(functions).ParseFiles(
			"./ui/html/pages/base.html",
			"./ui/html/partials/nav.html",
			page,
		)
		if err != nil {
			return err
		}

		// Add template to cache
		templateCache[name] = ts
	}

	return nil
}

func (app *application) render(w http.ResponseWriter, r *http.Request, page string, data templateData) {
	// Get template from cache
	ts, ok := templateCache[page]
	if !ok {
		app.serverError(w, r, fmt.Errorf("the template %s does not exist", page))
		return
	}

	// Execute the template
	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
