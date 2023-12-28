package domain

import "time"

type UserFactory struct{}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (f UserFactory) CreateNewUser(user *NewUser) *User {
	return &User{
		UserName:  user.UserName,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now().UTC(),
	}
}
