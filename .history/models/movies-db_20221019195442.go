package models

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id int) (*Movie, error) {
	return nil, nil
}

func (m *DBModel) GAll(id int) (*Movie, error) {
	return nil, nil
}