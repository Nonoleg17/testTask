package repo

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"testCase/internal/entity"
	"testCase/pkg/constants"
	"testCase/pkg/postgres"
)

type FriendshipRepo struct {
	pg *postgres.Postgres
}

func NewFriendshipRepo(pg *postgres.Postgres) *FriendshipRepo {
	return &FriendshipRepo{
		pg,
	}

}

func (fr *FriendshipRepo) FollowUser(ctx context.Context, follower string, target string) error {
	if follower == target {
		return errors.New("follower and target the same")
	}
	firstUser := &entity.Friendship{
		FirstUserId:  follower,
		SecondUserId: target,
	}
	secondUser := &entity.Friendship{
		FirstUserId:  target,
		SecondUserId: follower,
	}
	if err := fr.pg.DbConnect.First(firstUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

	} else {
		return errors.New("cannot follow the user again")
	}
	if err := fr.pg.DbConnect.First(secondUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		firstUser.FriendshipStatus = constants.Status1
		if err := fr.pg.DbConnect.Create(firstUser).Error; err != nil {
			return err
		}
	} else {
		if err := fr.pg.DbConnect.Model(secondUser).Update(&entity.Friendship{
			FriendshipStatus: constants.Status2,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}
