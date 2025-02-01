package models

import (
	"database/sql"
	"time"
)

type Event struct {
	ID          int
	Name        string
	Location    string
	Date        string
	Banner      string
	Thumbnail   string
	HomePage    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
	query := `SELECT id, name, location, date, banner, thumbnail, home_page, description, created_at, updated_at FROM events ORDER BY created_at DESC LIMIT 4`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Location, &event.Date, &event.Banner, &event.Thumbnail, &event.HomePage, &event.Description, &event.CreatedAt, &event.UpdatedAt)
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
