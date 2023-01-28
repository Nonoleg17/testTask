package usecase

import (
	"context"
)

type FriendshipUseCase struct {
	repo FriendshipRepo
}

func NewFriendshipUseCase(r FriendshipRepo) *FriendshipUseCase {
	return &FriendshipUseCase{r}
}

func (uc *FriendshipUseCase) FollowUser(ctx context.Context, follower string, target string) error {
	return uc.repo.FollowUser(ctx, follower, target)
}
