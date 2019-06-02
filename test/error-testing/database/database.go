package database

import "github.com/tommy-sho/golang-sandbox/test/error-testing/domain"

type Database interface {
	Create(user *domain.User) error
	GetByID(ID string) (*domain.User, error)
	UpdateUser(ID string, user *domain.User) error
}
