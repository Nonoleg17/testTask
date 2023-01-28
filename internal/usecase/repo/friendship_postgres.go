package repo

import (
	"context"
	"testCase/pkg/postgres"
)

type FriendshipRepo struct {
	*postgres.Postgres
}

func NewFriendshipRepo(pg *postgres.Postgres) *FriendshipRepo {
	return &FriendshipRepo{
		pg,
	}

}

func (ur *UserRepo) CreateFriendShip(ctx context.Context) error {
	return nil
}

func (ur *UserRepo) DeleteFriendShip(ctx context.Context) error {
	return nil
}

func (ur *UserRepo) UpdateFriendShip(ctx context.Context) error {
	return nil
}
