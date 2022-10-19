package models

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel)