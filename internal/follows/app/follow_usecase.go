package app

import (
	"example.com/my-medium-clone/internal/follows/domain"
)

type FollowUseCase interface {
	Create(follow *domain.Follow) (int, error)
	IsFollowing(follow *domain.Follow) (bool, error)
	Unfollow(follow *domain.Follow) error
}

type followUseCase struct {
	followRepo domain.FollowRepository
}

func NewFollowUseCase(followRepo domain.FollowRepository) FollowUseCase {
	return &followUseCase{followRepo: followRepo}
}

func (f *followUseCase) Create(follow *domain.Follow) (int, error) {
	isFollow, err := f.followRepo.IsFollowing(follow)
	if err != nil {
		return 0, err
	}

	if !isFollow {
		return f.followRepo.Save(follow)
	}
	return 0, err
}

func (f *followUseCase) IsFollowing(follow *domain.Follow) (bool, error) {
	return f.followRepo.IsFollowing(follow)
}

func (f *followUseCase) Unfollow(follow *domain.Follow) error {
	isFollow, err := f.followRepo.IsFollowing(follow)
	if err != nil {
		return err
	}

	if isFollow {
		return f.followRepo.Unfollow(follow)
	}

	return nil
}
