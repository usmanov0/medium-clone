package domain

type FollowRepository interface {
	Save(follow *Follow) (int, error)
	IsFollowing(follow *Follow) (bool, error)
	Unfollow(follow *Follow) error
}
