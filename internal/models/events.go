package models

import (
	"database/sql"
	"time"
)

type Event struct {
	ID          int
	Name        string
	Location    string
	DateStart   string
	DateEnd     string
	Banner      string
	Thumbnail   string
	HomePage    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	TalkCount   int
}

type EventModel struct {
	DB *sql.DB
}

func (m *EventModel) GetAll() ([]*Event, error) {
	return nil, nil
}

func (m *EventModel) Get(id int) (*Event, error) {
	return nil, nil

}

func (m *EventModel) Insert(event *Event) error {
	return nil
}

func (m *EventModel) Update(event *Event) error {
	return nil
}

func (m *EventModel) Delete(id int) error {
	return nil
}

func (m *EventModel) GetLatest() ([]Event, error) {
	query := `
		SELECT 
			e.id, e.name, e.location, e.date_start, e.date_end, 
			e.banner, e.thumbnail, e.home_page, e.description, 
			e.created_at, e.updated_at,
			COUNT(t.id) as talk_count
		FROM events e
		LEFT JOIN talks t ON e.id = t.event_id
		GROUP BY e.id, e.name, e.location, e.date_start, e.date_end,
				 e.banner, e.thumbnail, e.home_page, e.description,
				 e.created_at, e.updated_at
		ORDER BY e.created_at DESC 
		LIMIT 4`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID, &event.Name, &event.Location,
			&event.DateStart, &event.DateEnd, &event.Banner,
			&event.Thumbnail, &event.HomePage, &event.Description,
			&event.CreatedAt, &event.UpdatedAt, &event.TalkCount,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
