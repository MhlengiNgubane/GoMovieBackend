package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// Get returns one movie and error, if any
func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, year, release_date, runtime, mpaa_rating
			created_at, upd
	`
	return nil, nil
}

// All returns all movies and error, if any
func (m *DBModel) All(id int) ([]*Movie, error) {
	return nil, nil
}