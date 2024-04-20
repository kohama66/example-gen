package entity

import "time"

type User struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt time.Time
}

func NewUser(id int64, name, email string, createdAt time.Time) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedAt: createdAt,
	}
}
