package models

import "database/sql"

type DBModel struct {
	DB *sql.DB
}

// Get returns one movie and error, if any
func (m *DBModel) Get(id int) (*Movie, error) {
	return nil, nil
}

// All re
func (m *DBModel) All(id int) ([]*Movie, error) {
	return nil, nil
}