package adapters

import (
	"example.com/my-medium-clone/internal/users/domain"
	"example.com/my-medium-clone/internal/users/errors"
	"github.com/jackc/pgx"
)

func NewUserRepo(db *pgx.Conn) domain.UserRepository {
	return &userRepository{db: db}
}

type userRepository struct {
	db      *pgx.Conn
	userFac domain.UserFactory
}

func (u *userRepository) Save(user *domain.User) (id int, err error) {
	row := u.db.QueryRow("INSERT INTO users(user_name,email,password) VALUES($1,$2,$3) RETURNING id;",
		user.UserName, user.Email, user.Password)
	err = row.Scan(&id)
	if err != nil {
		return 0, errors.ErrIdScanFailed
	}
	return id, nil
}

func (u *userRepository) FindById(id int) (*domain.User, error) {
	row := u.db.QueryRow("SELECT u.id,u.user_name,u.email,u.bio,u.created_at,u.updated_at FROM users u WHERE b.id = $1",
		id)
	var user domain.User
	if err := row.Scan(user.Id, user.UserName, user.Email, user.Bio, user.CreatedAt, user.UpdatedAt); err != nil {
		return nil, err
	}
	return &user, nil

}

func (u *userRepository) ExistsByMail(email string) (bool, error) {
	var exists bool
	row := u.db.QueryRow("SELECT EXISTS(SELECT 1 from users WHERE email = $1)", email)
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *userRepository) Search(criteria string) ([]*domain.User, error) {
	query := "SELECT u.id,u.user_name,u.email,u.bio FROM users u WHERE user_name ILIKE $1 OR email ILIKE $1"
	rows, err := u.db.Query(query, "%"+criteria+"%")
	if err != nil {
		return nil, err
	}
	var users []*domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Bio); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) Update(userID int, user *domain.User) error {
	_, err := u.db.Exec("UPDATE users SET user_name = $1, password = $2, bio = $3 WHERE id = $4",
		user.UserName, user.Password, user.Bio, userID)
	if err != nil {
		return errors.ErrUserUpdateFailed
	}
	return nil
}

func (u *userRepository) Delete(id int) error {
	_, err := u.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return errors.ErrDeletingUser
	}
	return nil
}
