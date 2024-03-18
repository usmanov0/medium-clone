package impl

import (
	domain2 "example.com/my-medium-clone/internal/domain"
	"example.com/my-medium-clone/internal/errors"
	"example.com/my-medium-clone/internal/repo"
	"github.com/jackc/pgx"
)

type followRepository struct {
	db *pgx.Conn
}

func NewFollowRepository(db *pgx.Conn) repo.FollowRepository {
	return &followRepository{db: db}
}

func (f *followRepository) Save(follow *domain2.Follow) (int, error) {
	query := `
		INSERT INTO follows(followed_by_id) 
		VALUES($1) 
		RETURNING id`

	var id int
	err := f.db.QueryRow(query, follow.FollowedById).Scan(&id)
	if err != nil {
		return 0, errors.ErrIdScanFailed
	}
	return id, nil
}

func (f *followRepository) IsFollowing(follow *domain2.Follow) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM follows
			WHERE following_id = $1 and followed_by_id = $2)
			`

	var exists bool
	err := f.db.QueryRow(query, follow.FollowingId, follow.FollowedById).Scan(&exists)
	if err != nil {
		return false, errors.ErrIdScanFailed
	}
	return exists, nil
}

func (f *followRepository) Unfollow(follow *domain2.Follow) error {
	query := `
		DELETE FROM follows 
		WHERE following_id = $1 and followed_by_id = $2`

	_, err := f.db.Exec(query, follow.FollowedById, follow.FollowedById)
	if err != nil {
		return errors.ErrFailedExecuteQuery
	}
	return nil
}
