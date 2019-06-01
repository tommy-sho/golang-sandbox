package domain

import (
	"fmt"
	"time"
)

type User struct {
	ID        string
	Age       int
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u User) FullName() string {
	return fmt.Sprintf("%v-%v", u.FirstName, u.LastName)
}
