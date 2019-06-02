package domain

import (
	"fmt"
	"time"
)

type User struct {
	ID        string
	Age       int    `validate:"required,min=0"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u User) FullName() string {
	return fmt.Sprintf("%v-%v", u.FirstName, u.LastName)
}
