package entity

type UserRepository interface {
	Save(user *User) (int, error)
	FindById(userID int) (*User, error)
	ExistsByMail(email string) (bool, error)
}
