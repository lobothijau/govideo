package main

import (
	"errors"
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

func (app *application) talkDetail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/talks/"):])
	if err != nil || id < 1 {
		app.clientError(w, http.StatusNotFound)
		return
	}

	// Get the talk details
	talk, err := app.talks.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	// Get related talks from the same event
	relatedTalks, err := app.talks.GetRelatedTalks(talk.EventID, talk.ID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		Talk:         talk,
		RelatedTalks: relatedTalks,
	}

	app.render(w, r, "video_detail.html", data)
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

func (app *application) eventCreate(w http.ResponseWriter, r *http.Request) {
	data := templateData{}
	app.render(w, r, "event_create.html", data)
}

func (app *application) eventCreatePost(w http.ResponseWriter, r *http.Request) {
	// Limit the request body to 10MB
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Get form values
	event := &models.Event{
		Name:        r.PostFormValue("name"),
		Location:    r.PostFormValue("location"),
		DateStart:   r.PostFormValue("date_start"),
		DateEnd:     r.PostFormValue("date_end"),
		HomePage:    r.PostFormValue("home_page"),
		Description: r.PostFormValue("description"),
	}

	// Validate dates
	startDate, err := time.Parse("2006-01-02", event.DateStart)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", event.DateEnd)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	if endDate.Before(startDate) {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Handle banner image upload
	bannerFile, bannerHeader, err := r.FormFile("banner")
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	defer bannerFile.Close()

	// Generate unique filename for banner
	bannerFilename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), bannerHeader.Filename)
	event.Banner = bannerFilename

	// Create the banner file on disk
	bannerDst, err := os.Create(filepath.Join("./ui/static/images/events", bannerFilename))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	defer bannerDst.Close()

	// Copy the uploaded banner file
	_, err = io.Copy(bannerDst, bannerFile)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Handle thumbnail image upload
	thumbnailFile, thumbnailHeader, err := r.FormFile("thumbnail")
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	defer thumbnailFile.Close()

	// Generate unique filename for thumbnail
	thumbnailFilename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), thumbnailHeader.Filename)
	event.Thumbnail = thumbnailFilename

	// Create the thumbnail file on disk
	thumbnailDst, err := os.Create(filepath.Join("./ui/static/images/events", thumbnailFilename))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	defer thumbnailDst.Close()

	// Copy the uploaded thumbnail file
	_, err = io.Copy(thumbnailDst, thumbnailFile)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Insert the event into the database
	err = app.events.Insert(event)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Redirect to the home page (or event detail page once implemented)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) talkCreate(w http.ResponseWriter, r *http.Request) {
	// Get speakers for the dropdown
	speakers, err := app.speakers.GetAll()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Get events for the dropdown
	events, err := app.events.GetAll()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{}
	data.Speakers = speakers
	data.Events = events

	app.render(w, r, "talk_create.html", data)
}

func (app *application) talkCreatePost(w http.ResponseWriter, r *http.Request) {
	// Maximum file size of 10MB
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Get form values
	title := r.PostForm.Get("title")
	duration := r.PostForm.Get("duration")
	speakerID, err := strconv.Atoi(r.PostForm.Get("speaker_id"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	eventID, err := strconv.Atoi(r.PostForm.Get("event_id"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Handle thumbnail upload
	file, fileHeader, err := r.FormFile("thumbnail")
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate unique filename
	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), fileHeader.Filename)

	// Create the file
	dst, err := os.Create(filepath.Join("ui", "static", "images", "talks", filename))
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

	// Create talk model
	talk := &models.Talk{
		Title:     title,
		Duration:  duration,
		SpeakerID: speakerID,
		EventID:   eventID,
		Thumbnail: filename,
	}

	// Insert the talk
	_, err = app.talks.Insert(talk)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Redirect to the talk page
	http.Redirect(w, r, "/", http.StatusSeeOther)
	// http.Redirect(w, r, fmt.Sprintf("/talk/view/%d", id), http.StatusSeeOther)
}

func (app *application) talksView(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	perPage := 20
	offset := (page - 1) * perPage

	// Get talks for current page
	talks, err := app.talks.GetPaginated(offset, perPage)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Get total count of talks
	totalTalks, err := app.talks.GetTotalCount()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	totalPages := (totalTalks + perPage - 1) / perPage

	// Generate pagination numbers (show 5 pages around current page)
	var pages []int
	start := max(1, page-2)
	end := min(totalPages, page+2)

	for i := start; i <= end; i++ {
		pages = append(pages, i)
	}

	data := templateData{
		Talks:       talks,
		CurrentPage: page,
		TotalPages:  totalPages,
		HasNext:     page < totalPages,
		HasPrev:     page > 1,
		NextPage:    page + 1,
		PrevPage:    page - 1,
		Pages:       pages,
	}

	app.render(w, r, "talks.html", data)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (app *application) speakerDetail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/speakers/"):])
	if err != nil || id < 1 {
		app.clientError(w, http.StatusNotFound)
		return
	}

	speaker, err := app.speakers.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	talks, err := app.talks.GetBySpeaker(speaker.ID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		Speaker: speaker,
		Talks:   talks,
	}

	app.render(w, r, "speaker_detail.html", data)
}

func (app *application) speakersView(w http.ResponseWriter, r *http.Request) {
	speakers, err := app.speakers.GetAll()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		Speakers: speakers,
	}

	app.render(w, r, "speakers.html", data)
}
