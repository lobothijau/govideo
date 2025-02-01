package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", "4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /video", app.videoView)
	mux.HandleFunc("POST /video", app.videoCreate)
	mux.HandleFunc("GET /video/{id}", app.videoDetail)

	logger.Info("Starting server", "url", fmt.Sprintf("http://localhost:%s", *addr))

	port := fmt.Sprintf(":%s", *addr)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		logger.Error(err.Error())
	}
}
