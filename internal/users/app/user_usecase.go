package app

import (
	"example.com/my-medium-clone/internal/users/domain"
	"example.com/my-medium-clone/internal/users/errors"
)

type UserUseCase interface {
	SignUpUser(user *domain.NewUser) (int, error)
	SignInUser(email, password string) (bool, error)
}

type userUseCase struct {
	userRepo domain.UserRepository
	userFac  domain.UserFactory
}

func NewUserUseCase(userRepo domain.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (u userUseCase) SignUpUser(user *domain.NewUser) (int, error) {
	userFactory := u.userFac.CreateNewUser(user)

	err := validateUserInfoForSignUp(
		userFactory.UserName,
		userFactory.Email,
		userFactory.Password,
	)
	if err != nil {
		return 0, err
	}
	err = validateEmail(user.Email)
	if err != nil {
		return 0, errors.ErrInvalidEmailFormat
	}
	id, err := u.userRepo.Save(userFactory)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u userUseCase) SignInUser(email, password string) (bool, error) {
	err := validateUserInfoForSignIn(email, password)

	if err != nil {
		return false, err
	}
	exists, _ := u.userRepo.ExistsByMail(email)
	if !exists {
		return false, errors.ErrUserNotFound
	}
	return true, nil
}
