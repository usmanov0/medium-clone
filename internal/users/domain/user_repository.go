package domain

type UserRepository interface {
	Save(user *NewUserRepo) (int, error)
	SignIn(repo *SignInRepo) (string, error)
	GetFollowers(userId int) ([]*User, error)
	FindById(id int) (*User, error)
	FindOneByEmail(email string) (*User, error)
	ExistsByMail(email string) (bool, error)
	Search(criteria string) ([]User, error)
	Update(userID int, user *User) error
	Delete(id int) error
}
