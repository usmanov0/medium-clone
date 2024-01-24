package app

import (
	"example.com/my-medium-clone/internal/errors"
	"example.com/my-medium-clone/internal/users/domain"
	"fmt"
	"log"
)

type UserUseCase interface {
	SignUpUser(user *domain.NewUser) (int, error)
	SignInUser(email, password string) (bool, error)
	GetUserById(id int) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	ListUsers(criteria string) ([]domain.User, error)
	UpdateUser(userID int, user *domain.User) error
	DeleteUserAccount(id int) error
}

type userUseCase struct {
	userRepo domain.UserRepository
	userFac  domain.UserFactory
}

func NewUserUseCase(userRepo domain.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (u *userUseCase) SignUpUser(user *domain.NewUser) (int, error) {
	userFactory := u.userFac.CreateNewUser(user)

	err := domain.ValidateUserInfoForSignUp(
		userFactory.UserName,
		userFactory.Email,
		userFactory.Password,
	)
	log.Println(err)
	if err != nil {
		log.Println("Here")
		return 0, err
	}
	err = domain.ValidateEmail(user.Email)
	if err != nil {
		return 0, errors.ErrInvalidEmailFormat
	}
	id, err := u.userRepo.Save(userFactory)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userUseCase) SignInUser(email, password string) (bool, error) {
	err := domain.ValidateUserInfoForSignIn(email, password)

	if err != nil {
		return false, err
	}
	exists, _ := u.userRepo.ExistsByMail(email)
	if !exists {
		return false, errors.ErrUserNotFound
	}
	return true, nil
}

func (u *userUseCase) GetUserById(id int) (*domain.User, error) {
	user, err := u.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUseCase) GetUserByEmail(email string) (*domain.User, error) {
	user, err := u.userRepo.FindOneByEmail(email)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	return user, nil
}

func (u *userUseCase) ListUsers(criteria string) ([]domain.User, error) {
	userList, err := u.userRepo.Search(criteria)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve list of users %v", err)
	}

	return userList, nil
}

func (u *userUseCase) UpdateUser(userId int, userUpdate *domain.User) error {
	existingUser, err := u.userRepo.FindById(userId)
	if err != nil {
		return errors.ErrUserNotFound
	}

	existingUser.UserName = userUpdate.UserName
	existingUser.Password = userUpdate.Password
	existingUser.Bio = userUpdate.Bio

	err = u.userRepo.Update(userId, userUpdate)
	if err != nil {
		return errors.ErrUserUpdateFailed
	}
	return nil
}

func (u *userUseCase) DeleteUserAccount(id int) error {
	err := u.userRepo.Delete(id)
	if err != nil {
		return errors.ErrFailedDeleteAccount
	}
	return nil
}
