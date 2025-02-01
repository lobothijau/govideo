package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /video", app.videoView)
	mux.HandleFunc("POST /video", app.videoCreate)
	mux.HandleFunc("GET /video/{id}", app.videoDetail)

	return mux
}
