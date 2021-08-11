package user

import (
	"time"
)

type User struct {
	ID        int64
	UUID      string
	Email     string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
