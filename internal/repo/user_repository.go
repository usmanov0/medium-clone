package repo

import "example.com/my-medium-clone/internal/domain"

type UserRepository interface {
	Save(user *domain.User) (int, error)
	GetFollowers(userId int) ([]*domain.User, error)
	FindById(id int) (*domain.User, error)
	FindOneByEmail(email string) (*domain.User, error)
	ExistsByMail(email string) (bool, error)
	Search(criteria string) ([]domain.User, error)
	Update(userID int, user *domain.User) error
	Delete(id int) error
}
