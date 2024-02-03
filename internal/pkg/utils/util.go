package utils

import (
	"example.com/my-medium-clone/internal/errors"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"regexp"
	"strings"
)

func ValidateUserInfoForSignIn(email, password string) error {
	if strings.TrimSpace(email) == "" {
		return errors.ErrBadCredentials
	}
	if ValidatePassword(password) == errors.ErrInvalidPassword {
		return errors.ErrBadCredentials
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.ErrInvalidPassword
	}

	var (
		upperCase = regexp.MustCompile(`[A-Z]`)
		lowerCase = regexp.MustCompile(`[a-z]`)
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

func ValidateEmail(email string) error {
	emailReg := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailReg.MatchString(email) {
		return errors.ErrInvalidEmailFormat
	}
	return nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func RandomPassword() string {
	var digitBytes string = "0123456789"

	password := make([]byte, 4)
	for i := range password {
		password[i] = digitBytes[rand.Intn(len(digitBytes))]
	}

	return string(password)
}
