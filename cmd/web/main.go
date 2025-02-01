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

	logger.Info("Starting server", "url", fmt.Sprintf("http://localhost:%s", *addr))

	port := fmt.Sprintf(":%s", *addr)
	err := http.ListenAndServe(port, app.routes())
	if err != nil {
		logger.Error(err.Error())
	}
}
