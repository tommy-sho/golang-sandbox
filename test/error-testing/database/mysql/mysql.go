package mysql

import (
	"database/sql"

	"github.com/tommy-sho/golang-sandbox/test/error-testing/domain"
)

type database struct {
	client *sql.DB
}

func NewDatabase(client *sql.Conn) database.Database {
	return database{}
}

func (d database) Create(user *domain.User) error {
	return nil
}
func (d database) GetByID(ID string) (*domain.User, error) {
	return nil, nil
}
func (d database) UpdateUser(ID string, user *domain.User) error {
	return nil
}
