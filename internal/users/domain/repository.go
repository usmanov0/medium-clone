package domain

type UserRepository interface {
	Save(user *User) (int, error)
	FindById(id int) (*User, error)
	ExistsByMail(email string) (bool, error)
	Search(criteria string) ([]*User, error)
	Update(userID int, user *User) error
	Delete(id int) error
}
