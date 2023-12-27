package adapter

import (
	"example.com/my-medium-clone/internal/entity/user/entity"
	"github.com/jackc/pgx"
)

func NewUserRepo(db *pgx.Conn) entity.UserRepository {
	return &userRepository{db: db}
}

type userRepository struct {
	db      *pgx.Conn
	userFac entity.UserFactory
}

func (u userRepository) Save(user *entity.User) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindById(userID int) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) ExistsByMail(email string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
