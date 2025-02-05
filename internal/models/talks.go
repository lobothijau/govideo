package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Talk struct {
	ID            int
	Title         string
	Duration      string
	VideoID       string
	VideoProvider string
	SpeakerID     int
	Speaker       Speaker
	EventID       int
	Event         Event
	Thumbnail     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type TalkModel struct {
	DB *sql.DB
}

func (m *TalkModel) Insert(talk *Talk) (int, error) {
	query := `
		INSERT INTO talks (
			title, duration, speaker_id, event_id, thumbnail, video_id, video_provider
		)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	result, err := m.DB.Exec(query,
		talk.Title,
		talk.Duration,
		talk.SpeakerID,
		talk.EventID,
		talk.Thumbnail,
		talk.VideoID,
		talk.VideoProvider,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *TalkModel) GetLatest() ([]Talk, error) {
	query := `
		SELECT 
			t.id, t.title, t.duration, t.speaker_id, t.event_id, t.thumbnail, t.created_at, t.updated_at,
			s.id, s.name, s.created_at, s.updated_at,
			e.id, e.name, e.date_start, e.date_end, e.created_at, e.updated_at,
			t.video_id, t.video_provider
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
			&talk.VideoID, &talk.VideoProvider,
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

func (m *TalkModel) Get(id int) (*Talk, error) {
	query := `
		SELECT 
			t.id, t.title, t.duration, t.speaker_id, t.event_id, t.thumbnail, t.video_id, t.created_at, t.updated_at,
			s.id, s.name, s.avatar, s.created_at, s.updated_at,
			e.id, e.name, e.date_start, e.date_end, e.created_at, e.updated_at,
			t.video_id, t.video_provider
		FROM talks t
		LEFT JOIN speakers s ON t.speaker_id = s.id
		LEFT JOIN events e ON t.event_id = e.id
		WHERE t.id = ?
	`

	var talk Talk
	err := m.DB.QueryRow(query, id).Scan(
		&talk.ID, &talk.Title, &talk.Duration, &talk.SpeakerID, &talk.EventID,
		&talk.Thumbnail, &talk.VideoID, &talk.CreatedAt, &talk.UpdatedAt,
		&talk.Speaker.ID, &talk.Speaker.Name, &talk.Speaker.Avatar,
		&talk.Speaker.CreatedAt, &talk.Speaker.UpdatedAt,
		&talk.Event.ID, &talk.Event.Name, &talk.Event.DateStart,
		&talk.Event.DateEnd, &talk.Event.CreatedAt, &talk.Event.UpdatedAt,
		&talk.VideoID, &talk.VideoProvider,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return &talk, nil
}

func (m *TalkModel) GetRelatedTalks(eventID int, currentTalkID int) ([]Talk, error) {
	query := `
		SELECT 
			t.id, t.title, t.duration, t.speaker_id, t.event_id, t.thumbnail, t.created_at, t.updated_at,
			s.id, s.name, s.created_at, s.updated_at,
			t.video_id, t.video_provider
		FROM talks t
		LEFT JOIN speakers s ON t.speaker_id = s.id
		WHERE t.event_id = ? AND t.id != ?
		ORDER BY t.created_at DESC
		LIMIT 5
	`

	rows, err := m.DB.Query(query, eventID, currentTalkID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var talks []Talk

	for rows.Next() {
		var talk Talk
		err := rows.Scan(
			&talk.ID, &talk.Title, &talk.Duration, &talk.SpeakerID, &talk.EventID,
			&talk.Thumbnail, &talk.CreatedAt, &talk.UpdatedAt,
			&talk.Speaker.ID, &talk.Speaker.Name, &talk.Speaker.CreatedAt,
			&talk.Speaker.UpdatedAt,
			&talk.VideoID, &talk.VideoProvider,
		)
		if err != nil {
			return nil, err
		}
		talks = append(talks, talk)
	}

	return talks, nil
}

func (m *TalkModel) GetPaginated(offset, limit int) ([]Talk, error) {
	query := `
		SELECT 
			t.id, t.title, t.duration, t.speaker_id, t.event_id, t.thumbnail, 
			t.video_id, t.video_provider, t.created_at, t.updated_at,
			s.id, s.name, s.created_at, s.updated_at,
			e.id, e.name, e.date_start, e.date_end, e.created_at, e.updated_at
		FROM talks t
		LEFT JOIN speakers s ON t.speaker_id = s.id
		LEFT JOIN events e ON t.event_id = e.id
		ORDER BY t.created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := m.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var talks []Talk

	for rows.Next() {
		var talk Talk
		err := rows.Scan(
			&talk.ID, &talk.Title, &talk.Duration, &talk.SpeakerID, &talk.EventID,
			&talk.Thumbnail, &talk.VideoID, &talk.VideoProvider, &talk.CreatedAt, &talk.UpdatedAt,
			&talk.Speaker.ID, &talk.Speaker.Name, &talk.Speaker.CreatedAt, &talk.Speaker.UpdatedAt,
			&talk.Event.ID, &talk.Event.Name, &talk.Event.DateStart, &talk.Event.DateEnd,
			&talk.Event.CreatedAt, &talk.Event.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		talks = append(talks, talk)
	}

	return talks, nil
}

func (m *TalkModel) GetTotalCount() (int, error) {
	var count int
	err := m.DB.QueryRow("SELECT COUNT(*) FROM talks").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (m *TalkModel) GetBySpeaker(speakerID int) ([]Talk, error) {
	query := `
		SELECT 
			t.id, t.title, t.duration, t.speaker_id, t.event_id, t.thumbnail,
			t.video_id, t.video_provider, t.created_at, t.updated_at,
			e.id, e.name, e.date_start, e.date_end, e.created_at, e.updated_at
		FROM talks t
		LEFT JOIN events e ON t.event_id = e.id
		WHERE t.speaker_id = ?
		ORDER BY t.created_at DESC
	`

	rows, err := m.DB.Query(query, speakerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var talks []Talk

	for rows.Next() {
		var talk Talk
		err := rows.Scan(
			&talk.ID, &talk.Title, &talk.Duration, &talk.SpeakerID, &talk.EventID,
			&talk.Thumbnail, &talk.VideoID, &talk.VideoProvider, &talk.CreatedAt, &talk.UpdatedAt,
			&talk.Event.ID, &talk.Event.Name, &talk.Event.DateStart, &talk.Event.DateEnd,
			&talk.Event.CreatedAt, &talk.Event.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		talks = append(talks, talk)
	}

	return talks, nil
}

func (m *TalkModel) GetByEvent(eventID int) ([]Talk, error) {
	query := `
		SELECT 
			t.id, t.title, t.duration, t.thumbnail,
			t.speaker_id, t.event_id, t.created_at, t.updated_at,
			s.name as speaker_name, s.avatar as speaker_avatar
		FROM talks t
		JOIN speakers s ON t.speaker_id = s.id
		WHERE t.event_id = ?
		ORDER BY t.created_at ASC
	`

	rows, err := m.DB.Query(query, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var talks []Talk

	for rows.Next() {
		var talk Talk
		var speaker Speaker
		err := rows.Scan(
			&talk.ID, &talk.Title, &talk.Duration, &talk.Thumbnail,
			&talk.SpeakerID, &talk.EventID, &talk.CreatedAt, &talk.UpdatedAt,
			&speaker.Name, &speaker.Avatar,
		)
		if err != nil {
			return nil, err
		}
		talk.Speaker = speaker
		talks = append(talks, talk)
	}

	return talks, nil
}
