package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Speaker struct {
	ID        int
	Name      string
	Avatar    string
	HomePage  string
	Github    string
	Twitter   string
	Linkedin  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SpeakerModel struct {
	DB *sql.DB
}

func (m *SpeakerModel) GetAll() ([]*Speaker, error) {
	return nil, nil
}

func (m *SpeakerModel) Get(id int) (*Speaker, error) {
	return nil, nil
}

func (m *SpeakerModel) Insert(speaker *Speaker) error {
	return nil
}

func (m *SpeakerModel) Update(speaker *Speaker) error {
	return nil
}

func (m *SpeakerModel) Delete(id int) error {
	return nil
}

func (m *SpeakerModel) GetActive() ([]Speaker, error) {
	query := `
		SELECT * FROM speakers limit 6
	`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var speakers []Speaker

	for rows.Next() {
		var speaker Speaker
		err := rows.Scan(&speaker.ID, &speaker.Name, &speaker.Avatar, &speaker.HomePage, &speaker.Github, &speaker.Twitter, &speaker.Linkedin, &speaker.CreatedAt, &speaker.UpdatedAt)
		if err != nil {
			return nil, err
		}
		speakers = append(speakers, speaker)
	}

	for _, speaker := range speakers {
		fmt.Println(speaker)
	}

	return speakers, nil

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return speakers, nil
}
