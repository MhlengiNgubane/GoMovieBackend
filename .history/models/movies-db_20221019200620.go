package models

import "database/sql"

type DBModel struct {
	DB *sql.DB
}

// Get returns on
func (m *DBModel) Get(id int) (*Movie, error) {
	return nil, nil
}

func (m *DBModel) All(id int) ([]*Movie, error) {
	return nil, nil
}