package usecase

import (
	"context"
	uuid "github.com/satori/go.uuid"
)

type FriendshipUseCase struct {
	repo FriendshipRepo
}

func NewFriendshipUseCase(r FriendshipRepo) *FriendshipUseCase {
	return &FriendshipUseCase{r}
}

func (uc *FriendshipUseCase) FollowUser(ctx context.Context, follower uuid.UUID, target uuid.UUID) error {
	return uc.repo.FollowUser(ctx, follower, target)
}
