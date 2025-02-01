package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Talk struct {
	Title     string
	Duration  string
	Speaker   string
	Event     string
	Thumbnail string
}

type Speaker struct {
	Name      string
	TalkCount int
	Avatar    string
}

type Event struct {
	Name      string
	Location  string
	DateRange string
	TalkCount int
	Banner    string
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := templateData{}

	// Simulate talks data
	talks := []Talk{
		{
			Title:     "Building Fast Web Applications with Go",
			Duration:  "45m",
			Speaker:   "Jane Smith",
			Event:     "GopherCon 2025",
			Thumbnail: "https://placehold.co/600x400",
		},
		{
			Title:     "Advanced Database Patterns",
			Duration:  "30m",
			Speaker:   "John Doe",
			Event:     "GoWest 2025",
			Thumbnail: "https://placehold.co/600x400",
		},
		{
			Title:     "Microservices in Practice",
			Duration:  "60m",
			Speaker:   "Alice Johnson",
			Event:     "GoEurope 2025",
			Thumbnail: "https://placehold.co/600x400",
		},
		{
			Title:     "Security Best Practices",
			Duration:  "45m",
			Speaker:   "Bob Wilson",
			Event:     "SecureGo 2025",
			Thumbnail: "https://placehold.co/600x400",
		},
	}

	// Simulate speakers data
	speakers := []Speaker{
		{Name: "Jane Smith", TalkCount: 23, Avatar: "https://placehold.co/200x200"},
		{Name: "John Doe", TalkCount: 15, Avatar: "https://placehold.co/200x200"},
		{Name: "Alice Johnson", TalkCount: 31, Avatar: "https://placehold.co/200x200"},
		{Name: "Bob Wilson", TalkCount: 12, Avatar: "https://placehold.co/200x200"},
		{Name: "Carol White", TalkCount: 8, Avatar: "https://placehold.co/200x200"},
		{Name: "David Brown", TalkCount: 19, Avatar: "https://placehold.co/200x200"},
	}

	// Simulate events data
	events := []Event{
		{
			Name:      "GopherCon 2025",
			Location:  "San Francisco, USA",
			DateRange: "July 15-17, 2025",
			TalkCount: 45,
			Banner:    "https://placehold.co/800x400",
		},
		{
			Name:      "GoWest Summit",
			Location:  "Vancouver, Canada",
			DateRange: "August 10-12, 2025",
			TalkCount: 30,
			Banner:    "https://placehold.co/800x400",
		},
		{
			Name:      "GoEurope Conference",
			Location:  "Berlin, Germany",
			DateRange: "September 5-7, 2025",
			TalkCount: 25,
			Banner:    "https://placehold.co/800x400",
		},
	}

	data.Talks = talks
	data.Speakers = speakers
	data.Events = events

	app.render(w, r, http.StatusOK, "home.html", data)
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
