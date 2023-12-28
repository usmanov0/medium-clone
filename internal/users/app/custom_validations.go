package app

import (
	"example.com/my-medium-clone/internal/users/errors"
	"regexp"
	"strings"
)

func validateUserInfoForSignUp(userName, email, password string) error {

	if strings.TrimSpace(userName) == "" {
		return errors.ErrEmptyUserName
	}
	if strings.TrimSpace(email) == "" {
		return errors.ErrEmptyMail
	}
	if validatePassword(password) == errors.ErrInvalidPassword {
		return errors.ErrInvalidPassword
	}
	return nil
}

func validateUserInfoForSignIn(email, password string) error {
	if strings.TrimSpace(email) == "" {
		return errors.ErrBadCredentials
	}
	if validatePassword(password) == errors.ErrInvalidPassword {
		return errors.ErrBadCredentials
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.ErrInvalidPassword
	}

	var (
		upperCase = regexp.MustCompile(`[A-Z]`)
		lowerCase = regexp.MustCompile("`[a-z]`")
		digit     = regexp.MustCompile(`[0-9]`)
	)
	if !upperCase.MatchString(password) || !lowerCase.MatchString(password) {
		return errors.ErrInvalidPassword
	}
	if !digit.MatchString(password) {
		return errors.ErrInvalidPassword
	}
	return nil
}

func validateEmail(email string) error {
	emailReg := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailReg.MatchString(email) {
		return errors.ErrInvalidEmailFormat
	}
	return nil
}
