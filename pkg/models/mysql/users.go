package mysql

import (
	"database/sql"
	"nexttrack07/snippetbox/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

func (m *UserModel) Authenticate(email, password string) error {
	return nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
