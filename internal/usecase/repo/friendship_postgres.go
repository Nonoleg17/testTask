package repo

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
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

func (fr *FriendshipRepo) FollowUser(ctx context.Context, follower uuid.UUID, target uuid.UUID) error {
	if follower == target {
		return errors.New("follower and target the same")
	}
	friendship := &entity.Friendship{
		FirstUserId:  follower,
		SecondUserId: target,
	}

	if err := fr.pg.DbConnect.First(&entity.Friendship{}, "first_user_id = ? AND second_user_id = ?",
		follower, target).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

	} else {

		return errors.New("cannot follow the user again")
	}
	if err := fr.pg.DbConnect.First(&entity.Friendship{}, "first_user_id = ? AND second_user_id = ?",
		target, follower).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		friendship.FriendshipStatus = constants.Status1
		if err := fr.pg.DbConnect.Create(friendship).Error; err != nil {
			return err
		}
	} else {
		if err := fr.pg.DbConnect.Model(friendship).Update(&entity.Friendship{
			FriendshipStatus: constants.Status2,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}
