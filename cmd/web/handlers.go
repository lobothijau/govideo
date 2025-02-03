package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"govideo.dev/internal/models"
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

func (app *application) speakerCreate(w http.ResponseWriter, r *http.Request) {
	data := templateData{}

	app.render(w, r, "speaker_create.html", data)
}

func (app *application) speakerCreatePost(w http.ResponseWriter, r *http.Request) {
	// Limit the request body to 10MB
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Get form values
	speaker := &models.Speaker{
		Name:     r.PostFormValue("name"),
		HomePage: r.PostFormValue("home_page"),
		Github:   r.PostFormValue("github"),
		Twitter:  r.PostFormValue("twitter"),
		Linkedin: r.PostFormValue("linkedin"),
	}

	// Handle file upload
	file, fileHeader, err := r.FormFile("avatar")
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	defer file.Close()

	// Generate a unique filename for the avatar
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), fileHeader.Filename)
	speaker.Avatar = filename

	// Create the file on disk
	dst, err := os.Create(filepath.Join("./ui/static/images/speakers", filename))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the destination
	_, err = io.Copy(dst, file)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Insert the speaker into the database
	_, err = app.speakers.Insert(speaker)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Redirect to the home page (or speaker detail page once implemented)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
