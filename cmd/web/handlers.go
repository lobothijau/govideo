package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/pages/base.html",
		"./ui/html/pages/home.html",
		"./ui/html/partials/nav.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func videoView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Video View!"))
}

func videoDetail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/video/"):])
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Display video detail with ID  of %d", id)
	w.Write([]byte(msg))
}

func videoCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Video Create!"))
}
