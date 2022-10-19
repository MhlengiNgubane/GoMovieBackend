package models

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id int) (*m)