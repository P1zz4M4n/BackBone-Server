package domain

import (
	"context"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` //Not included in JSON responses
	Active    bool      `json:"active"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// UserRepository defines the interface for user-related database operations.
type UserRepository interface {
	createUser(ctx context.Context, user *User) error
	getUserByID(ctx context.Context, id string) (*User, error)
	getUserByEmail(ctx context.Context, email string) (*User, error)
	updateUser(ctx context.Context, user *User) error
	deleteUser(ctx context.Context, id string) error
}
