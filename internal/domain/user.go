package domain

import (
	"time"
)

type User struct {
	ID        string
	Name      string
	Surname   string
	Username  string
	Email     string
	Password  string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
