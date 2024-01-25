package adapters

import (
	"example.com/my-medium-clone/internal/errors"
	"example.com/my-medium-clone/internal/users/domain"
	"github.com/jackc/pgx"
)

type userRepository struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) domain.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Save(user *domain.User) (int, error) {
	query := `
		INSERT INTO users(user_name, email, password,bio)
		VALUES($1, $2, $3, $4)
		RETURNING id
	`

	var userID int
	err := u.db.QueryRow(query, user.UserName, user.Email, user.Password, user.Bio).Scan(&userID)
	if err != nil {
		return 0, errors.ErrIdScanFailed
	}
	return userID, nil
}

func (u *userRepository) GetFollowers(userId int) ([]*domain.User, error) {
	query := `
		SELECT u.id, u.user_name, u.email
		FROM users u
		INNER JOIN follows f ON u.id = f.following.id
		WHERE f.followed_by_id = $1`

	rows, err := u.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []*domain.User
	for rows.Next() {
		follower := &domain.User{}
		err := rows.Scan(&follower.Id, follower.UserName, follower.Email)
		if err != nil {
			return nil, errors.ErrIdScanFailed
		}
		followers = append(followers, follower)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return followers, nil
}

func (u *userRepository) FindById(id int) (*domain.User, error) {
	query := `
		SELECT id, user_name, email, password, bio, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user domain.User
	err := u.db.QueryRow(query, id).Scan(
		&user.Id, &user.UserName, &user.Email, &user.Password,
		&user.Bio, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (u *userRepository) FindOneByEmail(email string) (*domain.User, error) {
	var user domain.User

	query := `
		SELECT id,user_name,email,password,bio,created_at,updated_at	
		FROM users
		WHERE email = $1`

	err := u.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.UserName,
		&user.Email,
		&user.Password,
		&user.Bio,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err != nil {
		return nil, errors.ErrIdScanFailed
	}

	return &user, nil
}

func (u *userRepository) ExistsByMail(email string) (bool, error) {
	query := `
		SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)
	`

	var exists bool
	err := u.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, errors.ErrIdScanFailed
	}
	return exists, nil
}

func (u *userRepository) Search(criteria string) ([]domain.User, error) {
	query := `
		SELECT id, user_name, email, bio
		FROM users
		WHERE user_name ILIKE $1 OR email ILIKE $1
	`

	rows, err := u.db.Query(query, "%%"+criteria+"%%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Bio)
		if err != nil {
			return nil, errors.ErrIdScanFailed
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) Update(userID int, user *domain.User) error {
	query := `
		UPDATE users
		SET user_name = $1, password = $2, bio = $3, updated_at = $4
		WHERE id = $5
	`

	_, err := u.db.Exec(query, user.UserName, user.Password, user.Bio, user.UpdatedAt, userID)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Delete(id int) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := u.db.Exec(query, id)
	if err != nil {
		return errors.ErrFailedDeleteAccount
	}
	return nil
}
