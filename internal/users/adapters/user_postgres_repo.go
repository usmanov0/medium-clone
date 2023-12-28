package adapters

import (
	"example.com/my-medium-clone/internal/users/domain"
	"github.com/jackc/pgx"
)

func NewUserRepo(db *pgx.Conn) domain.UserRepository {
	return &userRepository{db: db}
}

type userRepository struct {
	db      *pgx.Conn
	userFac domain.UserFactory
}

func (u userRepository) Save(user *domain.User) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindById(id int) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) ExistsByMail(email string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
