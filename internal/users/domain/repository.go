package domain

type UserRepository interface {
	Save(user *User) (int, error)
	FindById(id int) (*User, error)
	ExistsByMail(email string) (bool, error)
}
