package usecase

import (
	domain2 "example.com/my-medium-clone/internal/domain"
	"example.com/my-medium-clone/internal/errors"
	"example.com/my-medium-clone/internal/repo"
	"fmt"
	"log"
	"time"
)

type UserUseCase interface {
	SignUpUser(user *domain2.NewUser) (int, error)
	SignInUser(email, password string) (bool, error)
	GetUserById(id int) (*domain2.User, error)
	GetUserByEmail(email string) (*domain2.User, error)
	ListUsers(criteria string) ([]domain2.User, error)
	UpdateUser(userID int, user *domain2.User) error
	DeleteUserAccount(id int) error
}

type userUseCase struct {
	userRepo repo.UserRepository
	userFac  domain2.UserFactory
}

func NewUserUseCase(userRepo repo.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (u *userUseCase) SignUpUser(user *domain2.NewUser) (int, error) {
	userFactory := u.userFac.CreateNewUser(user)

	err := domain2.ValidateUserInfoForSignUp(
		userFactory.UserName,
		userFactory.Email,
		userFactory.Password,
	)
	log.Println(err)
	if err != nil {
		log.Println("Here")
		return 0, err
	}
	err = domain2.ValidateEmail(user.Email)
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
	err := domain2.ValidateUserInfoForSignIn(email, password)

	if err != nil {
		return false, err
	}
	exists, _ := u.userRepo.ExistsByMail(email)
	if !exists {
		return false, errors.ErrUserNotFound
	}
	return true, nil
}

func (u *userUseCase) GetUserById(id int) (*domain2.User, error) {
	user, err := u.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUseCase) GetUserByEmail(email string) (*domain2.User, error) {
	user, err := u.userRepo.FindOneByEmail(email)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	return user, nil
}

func (u *userUseCase) ListUsers(criteria string) ([]domain2.User, error) {
	userList, err := u.userRepo.Search(criteria)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve list of users %v", err)
	}

	return userList, nil
}

func (u *userUseCase) UpdateUser(userId int, userUpdate *domain2.User) error {
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
