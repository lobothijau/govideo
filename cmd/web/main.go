package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"govideo.dev/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger   *slog.Logger
	events   *models.EventModel
	speakers *models.SpeakerModel
}

func main() {
	addr := flag.String("addr", "4000", "HTTP network address")

	dsn := flag.String("dsn", "root:@/govideo?parseTime=true", "MySQL data source name")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		logger:   logger,
		events:   &models.EventModel{DB: db},
		speakers: &models.SpeakerModel{DB: db},
	}
	logger.Info("Starting server", "url", fmt.Sprintf("http://localhost:%s", *addr))

	port := fmt.Sprintf(":%s", *addr)
	err = http.ListenAndServe(port, app.routes())
	if err != nil {
		logger.Error(err.Error())
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
