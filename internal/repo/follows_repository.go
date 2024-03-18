package repo

import "example.com/my-medium-clone/internal/domain"

type FollowRepository interface {
	Save(follow *domain.Follow) (int, error)
	IsFollowing(follow *domain.Follow) (bool, error)
	Unfollow(follow *domain.Follow) error
}
