package app

import (
	"example.com/my-medium-clone/internal/errors"
	"example.com/my-medium-clone/internal/pkg/utils"
	"example.com/my-medium-clone/internal/users/domain"
	"fmt"
	"time"
)

type UserUseCase interface {
	SignUpUser(user *domain.NewUser) (int, error)
	SignIn(user *domain.SignInUser) (bool, error)
	GetUserById(id int) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	ListUsers(criteria string) ([]domain.User, error)
	UpdateUser(userID int, user *domain.User) error
	DeleteUserAccount(id int) error
}

type userUseCase struct {
	userRepo domain.UserRepository
	userFac  *domain.UserFactory
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
		userFactory.Bio,
	)
	if err != nil {
		return 0, err
	}
	err = utils.ValidateEmail(user.Email)
	if err != nil {
		return 0, errors.ErrInvalidEmailFormat
	}

	id, err := u.userRepo.Save(userFactory)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *userUseCase) SignIn(user *domain.SignInUser) (bool, error) {
	err := utils.ValidateUserInfoForSignIn(user.Email, user.Password)
	if err != nil {
		return false, err
	}
	exists, _ := u.userRepo.ExistsByMail(user.Email)
	if !exists {
		return false, errors.ErrUserNotFound
	}

	resp := u.userFac.SignInEmailUser(user)
	hashedPassword, err := u.userRepo.SignIn(resp)
	if err != nil {
		return false, nil
	}
	err = utils.CheckPassword(resp.Password, hashedPassword)
	if err != nil {
		return false, nil
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

	if userUpdate.UserName != existingUser.UserName {
		existingUser.UserName = userUpdate.UserName
	} else {
		return errors.ErrShouldBeDifferentName
	}
	if userUpdate.Password != existingUser.Password {
		existingUser.Password = userUpdate.Password
	} else {
		return errors.ErrShouldBeDifferentPassword
	}
	if userUpdate.Bio != existingUser.Bio {
		existingUser.Bio = userUpdate.Bio
	} else {
		return errors.ErrShouldBeDifferentBio
	}

	existingUser.UpdatedAt = time.Now()

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
