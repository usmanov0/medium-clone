package domain

import "time"

type User struct {
	Id        int
	UserName  string
	Email     string
	Password  string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NewUser struct {
	UserName string
	Email    string
	Password string
}
