package models

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id int) (*Movie, error) {
	return nil, nil
}

func (m *DBModel) G(id int) (*Movie, error) {
	return nil, nil
}