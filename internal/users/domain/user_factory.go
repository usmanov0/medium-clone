package domain

import (
	"example.com/my-medium-clone/internal/errors"
	"example.com/my-medium-clone/internal/pkg/utils"
	"strings"
)

type UserFactory struct{}

func (f *UserFactory) CreateNewUser(user *NewUser) *NewUserRepo {

	hashPas, _ := utils.HashPassword(user.Password)
	return &NewUserRepo{
		UserName: user.UserName,
		Email:    user.Email,
		Password: hashPas,
		Bio:      user.Bio,
	}
}

func (f *UserFactory) SignInEmailUser(user *SignInUser) *SignInRepo {
	return &SignInRepo{
		Email:    user.Email,
		Password: user.Password,
	}
}

func ValidateUserInfoForSignUp(userName, email, password, _ string) error {

	if strings.TrimSpace(userName) == "" {
		return errors.ErrEmptyUserName
	}
	if strings.TrimSpace(email) == "" {
		return errors.ErrEmptyMail
	}
	if utils.ValidatePassword(password) == errors.ErrInvalidPassword {
		return errors.ErrInvalidPassword
	}
	return nil
}
