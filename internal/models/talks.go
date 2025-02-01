package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Talk struct {
	ID        int
	Title     string
	Duration  string
	SpeakerID int
	Speaker   Speaker
	EventID   int
	Event     Event
	Thumbnail string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TalkModel struct {
	DB *sql.DB
}

func (m *TalkModel) GetLatest() ([]Talk, error) {
	query := `
		SELECT 
			t.id, t.title, t.duration, t.speaker_id, t.event_id, t.thumbnail, t.created_at, t.updated_at,
			s.id, s.name, s.created_at, s.updated_at,
			e.id, e.name, e.date_start, e.date_end, e.created_at, e.updated_at
		FROM talks t
		LEFT JOIN speakers s ON t.speaker_id = s.id
		LEFT JOIN events e ON t.event_id = e.id
		ORDER BY t.created_at DESC LIMIT 4
	`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var talks []Talk

	for rows.Next() {
		var talk Talk
		err := rows.Scan(
			&talk.ID, &talk.Title, &talk.Duration, &talk.SpeakerID, &talk.EventID, &talk.Thumbnail, &talk.CreatedAt, &talk.UpdatedAt,
			&talk.Speaker.ID, &talk.Speaker.Name, &talk.Speaker.CreatedAt, &talk.Speaker.UpdatedAt,
			&talk.Event.ID, &talk.Event.Name, &talk.Event.DateStart, &talk.Event.DateEnd, &talk.Event.CreatedAt, &talk.Event.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		talks = append(talks, talk)
	}

	for _, talk := range talks {
		fmt.Println(talk)
	}

	return talks, nil
}
