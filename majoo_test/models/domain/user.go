package domain

import "time"

type User struct {
	ID        int64
	Name      string
	UserName  string
	Password  string
	CreatedAt time.Time
	CreatedBy int64
	UpdatedAt time.Time
	UpdatedBy int64
}
