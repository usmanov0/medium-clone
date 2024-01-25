package domain

type UserRepository interface {
	Save(user *User) (int, error)
	GetFollowers(userId int) ([]*User, error)
	FindById(id int) (*User, error)
	FindOneByEmail(email string) (*User, error)
	ExistsByMail(email string) (bool, error)
	Search(criteria string) ([]User, error)
	Update(userID int, user *User) error
	Delete(id int) error
}
