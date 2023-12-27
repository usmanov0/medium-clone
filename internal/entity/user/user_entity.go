package user

import "time"

type User struct {
	Id          int
	UserName    string
	Email       string
	Password    string
	Role        string
	Bio         string
	DateOfBirth time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserDto struct {
	UserName    string
	Email       string
	Password    string
	DateOfBirth time.Time
}
