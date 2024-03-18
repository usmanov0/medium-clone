package usecase

import (
	domain2 "example.com/my-medium-clone/internal/domain"
	"example.com/my-medium-clone/internal/repo"
)

type FollowUseCase interface {
	Create(follow *domain2.Follow) (int, error)
	IsFollowing(follow *domain2.Follow) (bool, error)
	Unfollow(follow *domain2.Follow) error
}

type followUseCase struct {
	followRepo repo.FollowRepository
}

func NewFollowUseCase(followRepo repo.FollowRepository) FollowUseCase {
	return &followUseCase{followRepo: followRepo}
}

func (f *followUseCase) Create(follow *domain2.Follow) (int, error) {
	isFollow, err := f.followRepo.IsFollowing(follow)
	if err != nil {
		return 0, err
	}

	if !isFollow {
		return f.followRepo.Save(follow)
	}
	return 0, err
}

func (f *followUseCase) IsFollowing(follow *domain2.Follow) (bool, error) {
	return f.followRepo.IsFollowing(follow)
}

func (f *followUseCase) Unfollow(follow *domain2.Follow) error {
	isFollow, err := f.followRepo.IsFollowing(follow)
	if err != nil {
		return err
	}

	if isFollow {
		return f.followRepo.Unfollow(follow)
	}

	return nil
}
