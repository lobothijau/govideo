package models

import (
	"database/sql"
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
	TalkCount int
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
		SELECT 
			s.id, s.name, s.avatar, s.home_page, s.github, s.twitter, s.linkedin, s.created_at, s.updated_at,
			COUNT(t.id) as talk_count
		FROM speakers s
		LEFT JOIN talks t ON s.id = t.speaker_id
		GROUP BY s.id, s.name, s.avatar, s.home_page, s.github, s.twitter, s.linkedin, s.created_at, s.updated_at
		ORDER BY s.created_at DESC
		LIMIT 6
	`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var speakers []Speaker

	for rows.Next() {
		var speaker Speaker
		err := rows.Scan(
			&speaker.ID, &speaker.Name, &speaker.Avatar, &speaker.HomePage, &speaker.Github, &speaker.Twitter, &speaker.Linkedin, &speaker.CreatedAt, &speaker.UpdatedAt, &speaker.TalkCount,
		)
		if err != nil {
			return nil, err
		}
		speakers = append(speakers, speaker)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return speakers, nil
}
